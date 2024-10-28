package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"errors"
	"sync/atomic"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// semaphore的实现 copy 自 shzf/bcp/pkg/semaphore/sem_darwin.go 这只是macOS下的实现，不同系统的实现不一样
// 按如下方法运行：
// 先启动reader： go run main.go -role reader
// 再启动writer： go run main.go -role writer -msg hello
//
// 基本思路是：需要两个信号量和一个mmap来实现单向的数据流动，即writer来写入完成后通知reader来读取，reader
// reader 读取完成后通知 writer 来写入，如此循环往复
// 若想实现双向通信，则需要4个信号量和2个mmap

// 参考：
// - https://www.youtube.com/watch?v=ukM_zzrIeXs
// - https://man7.org/linux/man-pages/man3/sem_open.3.html
// - https://www.cnblogs.com/dream397/p/14301620.html
// - https://www.cnblogs.com/zzhaolei/articles/17591442.html
// - https://github.com/cch123/golang-notes/blob/master/syscall.md

const (
	SharedFile       = "/tmp/shared_memory"
	ProducerSemaName = "producer"
	ConsumerSemaName = "consumer"
)

func openSharedMemory() (*os.File, []byte, error) {
	file, err := os.OpenFile(SharedFile, os.O_RDWR, 0666)
	if err != nil {
		return nil, nil, err
	}

	data, err := unix.Mmap(int(file.Fd()), 0, 1024, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return nil, nil, err
	}

	return file, data, nil
}

func createSharedMemory() (*os.File, []byte, error) {
	file, err := os.Create(SharedFile)
	if err != nil {
		return nil, nil, err
	}

	if err := file.Truncate(1024); err != nil {
		return nil, nil, err
	}

	data, err := unix.Mmap(int(file.Fd()), 0, 1024, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return nil, nil, err
	}

	return file, data, nil
}

func main() {
	role := flag.String("role", "reader", "the role of the game: reader/writer")
	msg := flag.String("msg", "", "msg to write")
	flag.Parse()

	// var file *os.File
	var sharedMemory []byte
	var err error

	if *role == "reader" {
		_, sharedMemory, err = createSharedMemory()
		if err != nil {
			fmt.Println("Error creating shared memory:", err)
			return
		}
		// defer file.Close()
	} else {
		_, sharedMemory, err = openSharedMemory()
		if err != nil {
			fmt.Println("Error opening shared memory:", err)
			return
		}
	}

	semProducer, err := NewPOSIXSem(ProducerSemaName, 0)
	if err != nil {
		panic(err)
	}
	// 初始值需要是1，否则reader和writer会死锁
	semConsumer, err := NewPOSIXSem(ConsumerSemaName, 1)
	if err != nil {
		panic(err)
	}

	if *role == "reader" {
		for {
			semProducer.Wait()
			fmt.Printf("Reading: %s\n", sharedMemory)
			semConsumer.Post()
		}
	} else {
		for i := 0; i < 5; i++ {
			semConsumer.Wait()
			copy(sharedMemory, *msg)
			semProducer.Post()
		}
	}

	// semProducer.Close()
	// semConsumer.Close()
}

const (
	// Semaphore default mode: 0777. default value: 0.
	semDefaultMode  uint32 = 0777
	semDefaultValue uint32 = 0
)

var (
	errNone = syscall.Errno(0)
	// ErrCanNotGet syscall SYS_SEM_GETVALUE macOS is not implemented.
	ErrCanNotGet = errors.New("can't get this value in macOS system")
)

// POSIXSem Linux posix semaphore.
type POSIXSem struct {
	name      string
	fd        uintptr
	mode      uint32
	initValue uint32
	wn        int32
}

// NewDefaultPOSIXSem request semaphore with default permission settings.
func NewDefaultPOSIXSem(name string) (*POSIXSem, error) {
	sem := new(POSIXSem)
	sem.name = name
	sem.mode = semDefaultMode
	sem.initValue = semDefaultValue
	p, err := syscall.BytePtrFromString(name)
	if err != nil {
		return nil, err
	}
	ret, err := semOpen(uintptr(unsafe.Pointer(p)), sem.mode, sem.initValue)
	if err != nil {
		return nil, err
	}
	sem.fd = ret
	return sem, nil
}

func NewPOSIXSem(name string, initVal uint32) (*POSIXSem, error) {
	sem := new(POSIXSem)
	sem.name = name
	sem.mode = semDefaultMode
	sem.initValue = initVal
	p, err := syscall.BytePtrFromString(name)
	if err != nil {
		return nil, err
	}
	ret, err := semOpen(uintptr(unsafe.Pointer(p)), sem.mode, sem.initValue)
	if err != nil {
		return nil, err
	}
	sem.fd = ret
	return sem, nil
}

func semOpen(namePtr uintptr, mode uint32, value uint32) (uintptr, error) {
	ret, _, err := syscall.Syscall6(syscall.SYS_SEM_OPEN, namePtr, syscall.O_CREAT|syscall.O_RDWR, uintptr(mode), uintptr(value), 0, 0)
	if err != errNone {
		return 0, err
	}
	return ret, nil
}

// DockerMode macOS don't care.
func (s *POSIXSem) DockerMode() error { return nil }

// Reset make sure the semaphore value is 0.
// In the easiest way, try to open the semaphore again after unlink.
func (s *POSIXSem) Reset() error {
	err := s.Close()
	if err != nil {
		return nil
	}
	err = s.Unlink()
	if err != nil {
		return err
	}
	ptr, err := s.getNamePtr()
	if err != nil {
		return err
	}
	fd, err := semOpen(ptr, s.mode, s.initValue)
	if err != nil {
		return err
	}
	s.fd = fd
	return nil
}

// Post increments (unlocks) the semaphore pointed to.
func (s *POSIXSem) Post() error {
	_, _, err := syscall.Syscall(syscall.SYS_SEM_POST, s.fd, 0, 0)
	if err != errNone {
		return err
	}
	return nil
}

// Wait decrements (locks) the semaphore pointed to.
func (s *POSIXSem) Wait() error {
	atomic.AddInt32(&s.wn, 1)
	var err error
	for {
		_, _, err = syscall.Syscall(syscall.SYS_SEM_WAIT, s.fd, 0, 0)
		if err == syscall.EINTR {
			continue
		}
		break
	}
	atomic.AddInt32(&s.wn, -1)
	if err != errNone {
		return err
	}
	return nil
}

// WaitWithEINTR decrements (locks) the semaphore pointed to.
// If an EINTR error occurs, it will not re-enter wait.
func (s *POSIXSem) WaitWithEINTR() error {
	atomic.AddInt32(&s.wn, 1)
	_, _, err := syscall.Syscall(syscall.SYS_SEM_WAIT, s.fd, 0, 0)
	atomic.AddInt32(&s.wn, -1)
	if err != errNone {
		return err
	}
	return nil
}

// GetValue get the current value of the semaphore pointed to.
// Calling on macOS will always return an error.
// macOS does not support this function.
func (s *POSIXSem) GetValue() (int, error) { return 0, ErrCanNotGet }

// Close closes the named semaphore referred to.
func (s *POSIXSem) Close() error {
	n := atomic.LoadInt32(&s.wn)
	var i int32
	if n > 0 {
		// Make sure there is no wait blocking.
		for ; i <= n; i++ {
			err := s.Post()
			if err != nil {
				return err
			}
		}
		for k := 0; k < 3; k++ {
			n := atomic.LoadInt32(&s.wn)
			if n == 0 {
				goto CLOSE
			}
			time.Sleep(2 * time.Millisecond)
		}
	}
CLOSE:
	_, _, err := syscall.Syscall(syscall.SYS_SEM_CLOSE, s.fd, 0, 0)
	if err != errNone {
		return err
	}
	return nil
}

func (s *POSIXSem) getNamePtr() (uintptr, error) {
	p, err := syscall.BytePtrFromString(s.name)
	return uintptr(unsafe.Pointer(p)), err
}

// Unlink removes the named semaphore referred to.
func (s *POSIXSem) Unlink() error {
	p, err := s.getNamePtr()
	if err != nil {
		return err
	}
	_, _, err = syscall.Syscall(syscall.SYS_SEM_UNLINK, p, 0, 0)
	if err != errNone {
		return err
	}
	return nil
}
