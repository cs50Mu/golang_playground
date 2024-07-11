package main

import (
	"fmt"
	"time"
)

// GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并
// 发安全，且需要实现以下接口：

// type sp interface {
//     Write(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
//     Read(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
// }

// ref: https://github.com/lifei6671/interview-go/blob/master/question/q010.md

func main() {
	bm := NewBlockingMap()

	go func() {
		time.Sleep(3 * time.Second)
		bm.Write("hello", "world")
		// fmt.Printf("after write: %+v\n", bm.innerM) // output for debug

	}()

	fmt.Printf("%v\n", bm.Read("hello", 5*time.Second))
}

type BlockingMap struct {
	m      map[string]chan struct{}
	innerM map[string]interface{}
}

func NewBlockingMap() *BlockingMap {
	return &BlockingMap{
		m:      make(map[string]chan struct{}),
		innerM: make(map[string]interface{}),
	}
}

// Read ...
func (bm *BlockingMap) Read(key string, timeout time.Duration) interface{} {
	if val, ok := bm.innerM[key]; ok {
		return val
	}

	var ch chan struct{}
	ch, ok := bm.m[key]
	if !ok {
		ch = make(chan struct{})
		bm.m[key] = ch
	}
	timer := time.After(timeout)
	select {
	case <-ch:
		// fmt.Printf("read: %+v\n", bm.innerM) // output for debug

		return bm.innerM[key]
	case <-timer:
		return nil
	}
}

// Write ...
func (bm *BlockingMap) Write(key string, val interface{}) {
	var ch chan struct{}
	ch, ok := bm.m[key]
	if !ok {
		ch = make(chan struct{})
		bm.m[key] = ch
	}
	bm.innerM[key] = val
	go func() {
		ch <- struct{}{}
	}()
}
