package main

import (
	"bufio"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

// 测试：
// 正向普通代理：curl -x http://localhost:19846/ -v http://www.baidu.com
// 正向tunnel代理：curl -x http://localhost:19846/ -v https://www.baidu.com
// curl 在使用代理的时候，当发现要请求的网站是https时，会自动转换为使用tunnel代理的方式来请求
// 学到一点，http代理同样可以代理https请求，比如可以设置 https_proxy=http://localhost:8964 只要这个http代理支持tunnel代理就行
// 参考：https://koalr.me/posts/passive-scan-via-http-proxy/

// 方法一：利用了net/http下的 http server ，比较优雅
// 但在当前的runtime（1.21.5）下，只有正向代理（handleHTTP）可以正常工作
// tunnel代理无法正常工作（curl提示Received HTTP code 404 from proxy after CONNECT）
// 通过单步调试找到原因了，通过 http.HandleFunc() 注册方式使用的是内置的DefaultServeMux，
// 而 CONNECT 请求的路径为空，跟注册的路径匹配不上，导致返回了404
// 改为自己实现 http handler 就没问题了

// 参考：https://eli.thegreenplace.net/2022/go-and-proxy-servers-part-2-https-proxies/
// https://koalr.me/posts/passive-scan-via-http-proxy/

type ProxyHttpServer struct {
	caCert *x509.Certificate
	caKey  any
	mitm   bool
}

// loadX509KeyPair loads a certificate/key pair from files, and unmarshals them
// into data structures from the x509 package. Note that private key types in Go
// don't have a shared named interface and use `any` (for backwards
// compatibility reasons).
func loadX509KeyPair(certFile, keyFile string) (cert *x509.Certificate, key any, err error) {
	cf, err := os.ReadFile(certFile)
	if err != nil {
		return nil, nil, err
	}

	kf, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, nil, err
	}
	certBlock, _ := pem.Decode(cf)
	cert, err = x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	keyBlock, _ := pem.Decode(kf)
	key, err = x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return cert, key, nil
}

// createCert creates a new certificate/private key pair for the given domains,
// signed by the parent/parentKey certificate. hoursValid is the duration of
// the new certificate's validity.
func createCert(dnsNames []string, parent *x509.Certificate, parentKey crypto.PrivateKey, hoursValid int) (cert []byte, priv []byte) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("Failed to generate serial number: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Sample MITM proxy"},
		},
		DNSNames:  dnsNames,
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Duration(hoursValid) * time.Hour),

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, parent, &privateKey.PublicKey, parentKey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}
	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if pemCert == nil {
		log.Fatal("failed to encode certificate to PEM")
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	if pemCert == nil {
		log.Fatal("failed to encode key to PEM")
	}

	return pemCert, pemKey
}

func (p *ProxyHttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("req method: %v", r.Method)
	if r.Method == http.MethodConnect {
		if p.mitm {
			p.handleMITMConnect(w, r)
		} else {
			p.handleConnect(w, r)
		}
	} else {
		p.handleHTTP(w, r)
	}
}

// 根证书可以使用 mkcert -install 来生成和安装到系统中
// 使用 mkcert -uninstall 来卸载
// 参考：https://github.com/FiloSottile/mkcert/issues/208
func main() {
	certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	port := flag.String("port", "19846", "the listening port of the proxy server")
	flag.Parse()

	portListen := fmt.Sprintf(":%v", *port)
	log.Printf("Starting proxy server on %v", portListen)

	handler := ProxyHttpServer{}
	cert, key, err := loadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Print("load cert and key failed, will not use mitm...")
	} else {
		log.Print("successful loaded cert and key, will use mitm")
		handler.caCert = cert
		handler.caKey = key
		handler.mitm = true
	}
	if err := http.ListenAndServe(portListen, &handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// 处理代理请求
func (p *ProxyHttpServer) handleHTTP(w http.ResponseWriter, r *http.Request) {
	// 解析目标 URL
	targetURL, err := url.Parse(r.RequestURI)
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusBadRequest)
		return
	}

	log.Printf("reqURL: %v\n", r.RequestURI)

	// 创建新的请求
	proxyReq, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// 复制请求头
	proxyReq.Header = r.Header
	log.Printf("request header: %+v", r.Header)

	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to reach target server", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// 复制响应头
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// 设置响应状态码
	w.WriteHeader(resp.StatusCode)
	log.Printf("resp header: %+v", resp.Header)
}

// proxyConnect implements the MITM proxy for CONNECT tunnels.
// 此种方案需要客户端信任一个特殊的自制根证书，可以实现中间人劫持，能看到
// 所有的请求和返回，虽然用的是https传输，平常用的抓包软件来抓取https流量
// 的底层原理就是这个
func (p *ProxyHttpServer) handleMITMConnect(w http.ResponseWriter, r *http.Request) {
	log.Printf("CONNECT requested to %v (from %v)", r.Host, r.RemoteAddr)

	// "Hijack" the client connection to get a TCP (or TLS) socket we can read
	// and write arbitrary data to/from.
	hj, ok := w.(http.Hijacker)
	if !ok {
		log.Fatal("http server doesn't support hijacking connection")
	}
	clientConn, _, err := hj.Hijack()
	if err != nil {
		log.Fatal("http hijack failed")
	}

	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		log.Fatal("error splitting host/port:", err)
	}

	// create a fake TLS certificate for the target host, signed by our CA
	pemCert, pemKey := createCert([]string{host}, p.caCert, p.caKey, 240)
	tlsCert, err := tls.X509KeyPair(pemCert, pemKey)
	if err != nil {
		log.Fatal(err)
	}

	// send http OK resp back to the client
	if _, err := clientConn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n")); err != nil {
		log.Fatal("error writing status to client:", err)
	}

	// create a new TLS server using our certificate. This server will now pretend
	// being the target
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		MinVersion:               tls.VersionTLS13,
		Certificates:             []tls.Certificate{tlsCert},
	}

	tlsConn := tls.Server(clientConn, tlsConfig)
	defer tlsConn.Close()

	// create a buffered reader for the client conn; this is required to use
	// http package function with is conn
	connReader := bufio.NewReader(tlsConn)

	// run the proxy in a loop until the client closes the conn
	for {
		// read an http request from the client; the request is sent over TLS that
		// connReader is configured to serve. The read will run a TLS handshake in
		// the first invocation(we could also call tlsConn.Handshake explicitly before the loop,
		// but this isn't necessary).
		// note that while the client believes it's talking across an encrypted channel
		// with the target, the proxy gets these requests in "plain text" because of the MITM setup
		req, err := http.ReadRequest(connReader)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// we can dump the request; log it or modify it, etc..
		if b, err := httputil.DumpRequest(req, false); err == nil {
			log.Printf("incoming request:\n%s\n", string(b))
		}

		// Take the original request and changes its destination to be forwarded
		// to the target server
		changeRequestToTarget(req, r.Host)

		// send the request to the target server
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal("err sending req to target:", err)
		}
		if b, err := httputil.DumpResponse(resp, false); err == nil {
			log.Printf("target response:\n%s\n", string(b))
		}
		defer resp.Body.Close()

		// send the target server's response back to the client
		if err := resp.Write(tlsConn); err != nil {
			log.Println("err writing response back to the client:", err)
		}
	}
}

// changeRequestToTarget modifies req to be re-routed to the given target;
// the target should be taken from the Host of the original tunnel (CONNECT)
// request.
func changeRequestToTarget(req *http.Request, targetHost string) {
	targetUrl := addrToUrl(targetHost)
	targetUrl.Path = req.URL.Path
	targetUrl.RawQuery = req.URL.RawQuery
	req.URL = targetUrl
	// Make sure this is unset for sending the request through a client
	req.RequestURI = ""
}

func addrToUrl(addr string) *url.URL {
	if !strings.HasPrefix(addr, "https") {
		addr = "https://" + addr
	}
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

// 处理 CONNECT 方法
// 这是普通的隧道代理，可以代理http和https，但只是原样转发，因此并不能解析出https的原始请求和返回内容
func (p *ProxyHttpServer) handleConnect(w http.ResponseWriter, r *http.Request) {
	// 解析目标地址
	host := r.Host
	log.Printf("connect host: %v\n", host)

	// 与目标地址建立连接
	destConn, err := net.Dial("tcp", host)
	if err != nil {
		http.Error(w, "Failed to connect to target server", http.StatusBadGateway)
		return
	}
	defer destConn.Close()

	// 向客户端发送连接成功的响应
	w.WriteHeader(http.StatusOK)
	h, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := h.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	// 将客户端连接和目标连接进行双向复制
	go io.Copy(destConn, clientConn)
	io.Copy(clientConn, destConn)
}

var client = http.Client{}

//===================================================================
// 方法二：这种方法没有利用net/http里的http server，而是自己
// 创建了一个tcp socket，相对于方法一更底层。
// 此方法下两种代理都可以正常运行

func main2() {
	listener, err := net.Listen("tcp", "127.0.0.1:19846")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	// 读取代理中的请求
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Println(err)
		return
	}
	req.RequestURI = ""
	log.Printf("req method: %v\n", req.Method)
	if req.Method == http.MethodConnect {
		// 解析目标地址
		host := req.Host
		log.Printf("connect host: %v\n", host)

		// 与目标地址建立连接
		destConn, err := net.Dial("tcp", host)
		if err != nil {
			log.Println(err)
			return
		}
		defer destConn.Close()
		log.Printf("destConn: %+v\n", destConn)

		// 向客户端发送连接成功的响应
		resp := new(http.Response)
		resp.Status = "Connection Established"
		resp.StatusCode = 200
		resp.Proto = "HTTP/1.1"
		resp.ProtoMajor = 1
		resp.ProtoMinor = 1
		resp.Write(conn)

		// 将客户端连接和目标连接进行双向复制
		go io.Copy(destConn, conn)
		io.Copy(conn, destConn)
	} else {
		// 发送请求获取响应
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		// 将响应返还给客户端
		_ = resp.Write(conn)
		_ = conn.Close()
	}
}
