package main

import (
	"fmt"
	"sync"
)

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

// func main() {
// 	// fmt.Println("vim-go")

// 	// counter := 1
// 	letterChan := make(chan rune)
// 	numChan := make(chan int)

// 	go func() {
// 		letter := 'A'
// 		for letter <= 'Z' {
// 			// fmt.Printf("%c\n", letter)
// 			letterChan <- letter
// 			letter += 1
// 		}
// 		close(letterChan)
// 	}()

// 	go func() {
// 		for i := 1; i < 100; i++ {
// 			numChan <- i
// 		}
// 	}()

// 	exit := false
// 	for {
// 		fmt.Printf("%v", <-numChan)
// 		fmt.Printf("%v", <-numChan)
// 		if exit {
// 			break
// 		}
// 		fmt.Printf("%c", <-letterChan)
// 		letter := <-letterChan
// 		fmt.Printf("%c", letter)
// 		if letter == 'Z' {
// 			exit = true
// 		}
// 	}
// }

func main() {
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
