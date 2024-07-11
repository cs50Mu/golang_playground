package main

import (
	"sync"
	"time"
)

func main() {
	gp := New()
	gp.Start(5)

	taskFunc := func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	for i := 0; i < 50; i++ {
		gp.AddTask(taskFunc)
	}

	gp.StopAndWait()
}

type gpool struct {
	taskQ chan Task
	// workerM map[int]*Worker
	workers     []*Worker
	workerStack []int
	m           sync.Mutex
	cond        sync.Cond
}

type Worker struct {
	id    int
	taskQ chan Task
}

type Task func() error

func newWorker(idx int) *Worker {
	return &Worker{
		id:    idx,
		taskQ: make(chan Task),
	}
}

func (w *Worker) run(pool *gpool, workerIdx int) {
	for task := range w.taskQ {
		// fmt.Printf("worker[%v] runing task-%v\n", w.id, task.id)
		task()
		pool.pushWorker(workerIdx)
	}
}

func New() *gpool {
	gp := &gpool{
		taskQ: make(chan Task),
	}
	gp.cond = *sync.NewCond(&gp.m)
	return gp
}

func (gp *gpool) Start(n int) {
	gp.workers = make([]*Worker, n)
	gp.workerStack = make([]int, n)
	for i := 0; i < n; i++ {
		worker := newWorker(i)
		// gp.workerM[i] = &worker
		gp.workers[i] = worker
		gp.workerStack[i] = i
		go worker.run(gp, i)
	}

	go gp.dispatch()
}

func (gp *gpool) dispatch() {
	for task := range gp.taskQ {
		gp.cond.L.Lock()
		for len(gp.workerStack) == 0 {
			gp.cond.Wait()
		}
		gp.cond.L.Unlock()

		worker := gp.popWorker()
		worker.taskQ <- task
	}
}

func (gp *gpool) popWorker() *Worker {
	gp.m.Lock()
	defer gp.m.Unlock()

	workerIdx := gp.workerStack[len(gp.workerStack)-1]
	gp.workerStack = gp.workerStack[:len(gp.workerStack)-1]
	return gp.workers[workerIdx]
}

func (gp *gpool) pushWorker(idx int) {
	gp.m.Lock()
	gp.workerStack = append(gp.workerStack, idx)
	gp.m.Unlock()

	gp.cond.Signal()
}

func (gp *gpool) AddTask(task Task) {
	// add task to the task pool
	gp.taskQ <- task
}

// StopAndWait stop the pool and wait for the submitted tasks to complete
func (gp *gpool) StopAndWait() {
	for {
		gp.m.Lock()
		// gp.workers 是否也应该加锁获取？
		workerStackLen := len(gp.workerStack)
		workerLen := len(gp.workers)
		gp.m.Unlock()

		// 没有新任务了而且也没有worker在工作了
		// 证明可以安全退出了
		if len(gp.taskQ) == 0 && workerStackLen == workerLen {
			break
		}

		time.Sleep(50 * time.Millisecond)
	}
}
