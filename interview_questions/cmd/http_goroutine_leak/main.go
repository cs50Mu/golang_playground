package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
)

// 一次请求会对应两个协程：一个读，一个写
// 若不对返回的response做任何处理（不读也不关），则会泄漏两个协程

func main() {
	num := 6
	for index := 0; index < num; index++ {
		resp, _ := http.Get("https://www.baidu.com")
		// http.Get("https://www.baidu.com")
		_, _ = io.ReadAll(resp.Body)
		resp.Body.Close()
	}
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}
