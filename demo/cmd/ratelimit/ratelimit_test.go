package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCounterLimiter(t *testing.T) {

	// cl := NewCounterLimiter(5*time.Second, 30)
	// cl := NewLeakBucketLimiter(5, 3)
	cl := NewTokenBucketLimiter(5, 3)

	turns := 5
	workersCnt := 2
	var limited atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < workersCnt; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < turns; i++ {
				res := cl.Allow()
				if res {
					fmt.Println("pass")
				} else {
					limited.Add(1)
					fmt.Println("not pass")
				}
				time.Sleep(200 * time.Millisecond)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("限制的次数为: %v\n通过的次数为: %v\n限制的比例为: %v\n",
		limited.Load(),
		workersCnt*turns-int(limited.Load()),
		float64(limited.Load())/float64(workersCnt*turns))
}
