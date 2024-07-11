package main

import (
	"testing"
	"time"
)

const (
	// PoolSize = 1e4
	PoolSize = 2e5
	TaskNum  = 1e6
)

func BenchmarkGoPool(b *testing.B) {
	taskFunc := func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := NewPool(PoolSize)
		for num := 0; num < TaskNum; num++ {
			p.AddTask(taskFunc)
		}
		p.StopAndWait()
	}
	b.StopTimer()
}
