package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排序。限时5秒，使用多个goroutine查找切片中是否存在给定的值，在查找到目标值或者超时后立刻结束所有goroutine的执行。

// 比如，切片 [23,32,78,43,76,65,345,762,......915,86]，查找目标值为 345 ，如果切片中存在，则目标值输出"Found it!"并立即取消仍在执行查询任务的goroutine。

// 如果在超时时间未查到目标值程序，则输出"Timeout！Not Found"，同时立即取消仍在执行的查找任务的goroutine。

// ref: https://github.com/lifei6671/interview-go/blob/master/question/q017.md

func main() {
	// prepare nums
	var nums []int
	total := 100000000
	for i := 0; i < total; i++ {
		nums = append(nums, rand.Int())
	}
	// decide target
	target := nums[rand.Intn(total)]

	// find target
	ctx, cancel := context.WithCancel(context.Background())
	step := 10000000
	var wg sync.WaitGroup
	doneChan := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneChan)
	}()

	for i := 0; i <= total-step; i += step {
		high := i + step
		if high > total {
			high = total
		}
		wg.Add(1)
		go worker(ctx, nums[i:high], target, cancel, &wg)
	}

	timeout := 1 * time.Second
	// timeout := 5 * time.Millisecond
	timer := time.NewTimer(timeout)
	select {
	case <-timer.C:
		cancel()
		fmt.Println("Timeout, not found!")
		return
	case <-doneChan:
		fmt.Println("Found it!")
		return
	}
}

func worker(ctx context.Context, nums []int, target int,
	cancelFunc context.CancelFunc,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for _, num := range nums {
		select {
		case <-ctx.Done():
			return
		default:
			if num == target {
				cancelFunc()
				return
			}
		}
	}
}
