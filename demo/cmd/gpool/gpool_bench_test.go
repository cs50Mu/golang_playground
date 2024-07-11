package main

import (
	"testing"
	"time"
)

const (
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
		gp := New()
		gp.Start(PoolSize)
		for i := 0; i < TaskNum; i++ {
			gp.AddTask(taskFunc)
		}
		gp.StopAndWait()
	}
	b.StopTimer()
}
