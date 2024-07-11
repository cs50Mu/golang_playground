package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	p := NewPool(200000)

	for i := 0; i < 1000000; i++ {
		idx := i
		p.AddTask(func() error {
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("processing task: %v\n", idx)
			return nil
		})
	}

	p.StopAndWait()
}

type Task func() error

type Pool struct {
	Closed        atomic.Bool
	taskWg        sync.WaitGroup
	taskCloseOnce sync.Once
	workerWg      sync.WaitGroup
	taskChan      chan Task
	poolSize      int
	ctx           context.Context
	cancelFunc    context.CancelFunc
}

func NewPool(n int) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	p := &Pool{
		poolSize: n,
		taskChan: make(chan Task, 1000000),
		// taskChan:   make(chan Task),
		ctx:        ctx,
		cancelFunc: cancel,
	}
	p.Start()

	return p
}

func (p *Pool) Start() {
	for i := 0; i < p.poolSize; i++ {
		p.workerWg.Add(1)
		go p.worker()
	}
}

func (p *Pool) AddTask(t Task) {
	closed := p.Closed.Load()
	if closed {
		return
	}

	p.taskWg.Add(1)
	p.taskChan <- t
}

func (p *Pool) StopAndWait() {
	p.Closed.Store(true)

	// wait for remaining tasks to be processed
	p.taskWg.Wait()

	// close task channel
	p.taskCloseOnce.Do(func() {
		close(p.taskChan)
	})

	// wait for workers to exit
	p.cancelFunc()
	p.workerWg.Wait()
}

func (p *Pool) worker() {
	defer func() {
		p.workerWg.Done()
	}()

	for {
		select {
		case <-p.ctx.Done():
			return
		// https://stackoverflow.com/questions/46200343/force-priority-of-go-select-statement
		// Prioritize context.Done statement
		case task, ok := <-p.taskChan:
			select {
			case <-p.ctx.Done():
				// 若有任务，需要把它标记完成，否则taskWg.Wait()会一直等待
				if task != nil && ok {
					p.taskWg.Done()
				}
				return
			default:
				if task == nil || !ok {
					return
				}
				task()
				p.taskWg.Done()
			}
		}
	}
}
