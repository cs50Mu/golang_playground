package main

import (
	"fmt"
	"time"
)

// 参考：https://www.youtube.com/watch?v=5zXAHh5tJqQ

type Task func() error

type Token struct{}

type pool struct {
	TaskChan chan Task
	Limit    int
	Sem      chan Token
}

func NewPool(limit int) *pool {
	p := &pool{
		// taskChan 必须是一个unbuffered channel才行
		TaskChan: make(chan Task),
		Limit:    limit,
		Sem:      make(chan Token, limit),
	}
	go p.Start()

	return p
}

func (p *pool) Start() {
	for task := range p.TaskChan {
		t := task
		// 限制并发度
		p.Sem <- Token{}
		// 这种方法的一个弊端是需要不断地启动新的协程
		// 但go的runtime应该是会有协程复用的，开销并不大？
		// 实际bench了一下，看起来内存方面的占用并不好，比复用goroutine要差不少
		go func() {
			t()
			<-p.Sem
		}()
	}
}

func (p *pool) AddTask(t Task) {
	p.TaskChan <- t
}

func (p *pool) StopAndWait() {
	close(p.TaskChan)

	// The loop is there to wait for the completion of the =last worker goroutine=.
	// Once the last item of hugeSlice has been read, the last goroutine is started
	// and the main goroutine goes on immediately since the first loop has ended.
	for n := p.Limit; n > 0; n-- {
		p.Sem <- Token{}
	}
}

func main() {
	p := NewPool(200000)
	// go p.Start()
	taskFunc := func() error {
		fmt.Println("xxx")
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	for i := 0; i < 1000000; i++ {
		p.AddTask(taskFunc)
	}

	p.StopAndWait()
}
