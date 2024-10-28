package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func letterNumAlt() {
	letterChan := make(chan struct{})
	digitChan := make(chan struct{})
	doneChan := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 1
		running := true
		for running {
			fmt.Printf("%v", i)
			fmt.Printf("%v", i+1)
			select {
			case <-doneChan:
				running = false
			default:
				digitChan <- struct{}{}
				i += 2
				<-letterChan
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		letter := 'A'
		for letter <= 'Z' {
			<-digitChan
			fmt.Printf("%c", letter)
			fmt.Printf("%c", letter+1)
			letter += 2
			letterChan <- struct{}{}
		}
		close(doneChan)
	}()

	wg.Wait()
}

func main() {
	// letterNumAlt()
	// altOddEven()

	// altOddEven2()

	// seqPrint(5)
	// seqPrint2(5)
	// orderPrint()

	// deadLockDemo()
	// gRet()

	// plusFunc := closureDemo()
	// plusFunc()
	// plusFunc()
	// plusFunc()
	// ctxDemo()
	// shallowCopyDeepCopy()
	panicRecoverDemo()
}

// 使用两个goroutine交替打印1-100之间的奇数和偶数, 输出时按照从小到大输出.
func altOddEven() {
	oddChan := make(chan struct{})
	evenChan := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i += 2 {
			fmt.Printf("%v", i)
			oddChan <- struct{}{}
			<-evenChan
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-oddChan
			fmt.Printf("%v", i)
			evenChan <- struct{}{}
		}
	}()

	wg.Wait()
}

// 创建N个协程，每个协程负责打印一个数字，编程实现将所有数字顺次输出。
// Input:
//     N = 5
// Output:
//     1
//     2
//     3
//     4
//     5
// ref: https://segmentfault.com/a/1190000042750846

// 使用锁
func seqPrint(n int) {
	num := 1
	var m sync.Mutex
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				m.Lock()
				// check if it's its turn to print
				if id == num {
					fmt.Println(num)
					num++
					m.Unlock()
					break
				} else {
					// if not, release the lock
					// to let others to print
					m.Unlock()
				}
			}
		}(i)
	}

	wg.Wait()
}

// 使用 N 个 channel, 每个 goroutine 阻塞等待上一个
// goroutine 完成，才输出，输出完成后再发信号告诉下一个
// goroutine 来继续
func seqPrint2(n int) {
	chans := make([]chan bool, n+1)
	for i := 0; i <= n; i++ {
		chans[i] = make(chan bool)
	}

	for i := 1; i <= n; i++ {
		go func(i int) {
			// wait for the previous goroutine to finish
			<-chans[i-1]
			fmt.Println(i)
			// signal to next goroutine
			chans[i] <- true
		}(i)
	}

	// signal the first goroutine
	chans[0] <- true

	// wait for the last goroutine to finish
	<-chans[n]
}

// 现在有4个协程，分别对应编号为1,2,3,4,每秒钟就有一个协程打印自己的编
// 号，要求编写一个程序，让输出的编号总是按照1,2,3,4,1,2,3,4这样的规律
// 一直打印下去

// 1 wait for chans[0]
// 2 wait for chans[1]
// 3 wait for chans[2]
// 4 wait for chans[3] and signals chans[0]
func orderPrint() {
	chans := make([]chan bool, 4)
	for i := 0; i < 4; i++ {
		chans[i] = make(chan bool)
	}
	for i := 0; i <= 3; i++ {
		go func(i int, previousChan, nextChan chan bool) {
			for {
				<-previousChan
				fmt.Println(i + 1)
				time.Sleep(1 * time.Second)
				nextChan <- true
			}
		}(i, chans[i], chans[(i+1)%4])
	}

	chans[0] <- true

	select {}
}

// 以下题目来自此篇帖子
// ref: https://juejin.cn/post/7126020110854127653
// 实现单例模式
type Config struct{}

var globalConf *Config
var m sync.Mutex

func GetConf() *Config {
	m.Lock()
	defer m.Unlock()

	if globalConf == nil {
		globalConf = &Config{}
	}
	return globalConf
}

// 写一个死锁
func deadLockDemo() {
	var m1 sync.Mutex
	var m2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()

		m1.Lock()
		time.Sleep(50 * time.Millisecond)
		defer m1.Unlock()

		m2.Lock()
		defer m2.Unlock()
	}()

	go func() {
		defer wg.Done()

		m2.Lock()
		time.Sleep(100 * time.Millisecond)
		defer m2.Unlock()

		m1.Lock()
		defer m1.Unlock()
	}()

	wg.Wait()
}

// 如何从 go routine 中获取返回值
// gRet ...
func gRet() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		val := <-ch
		fmt.Printf("got val in g: %v\n", val)
	}()

	fmt.Printf("send val from main\n")
	ch <- 1

	wg.Wait()
}

// 写一个闭包
// closureDemo ...
func closureDemo() func() {
	sum := 0

	plusOne := func() {
		sum += 1
		fmt.Printf("sum: %v\n", sum)
	}

	return plusOne
}

// Context 的用法
// ctxDemo ...
func ctxDemo() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("not time\n")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	fmt.Printf("time up, bye!\n")
	cancelFunc()
}

// 实现 slice 的深拷贝 和 浅拷贝
// shallowCopyDeepCopy ...
func shallowCopyDeepCopy() {
	a := []int{1, 2, 3}
	b := make([]int, 3)
	copy(b, a)
	a[1] = 100
	fmt.Printf("a: %v, b: %v\n", a, b)

	c := []int{4, 5, 6}
	d := c
	c[1] = -4
	fmt.Printf("c: %v, d: %v\n", c, d)
}

// pannic 和 recover
// panicRecoverDemo ...
func panicRecoverDemo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("catched panic err: %v\n", err)
		}
	}()

	panic("the world is going to explode!")
}
