package main

import (
	"fmt"
	"time"
)

type Item int

type Queue struct {
	items chan []Item
	empty chan bool
}

func main() {
	q := NewQueue()

	go func() {
		for i := 0; i < 10; i++ {
			q.Put(Item(i))
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			item := q.Get()
			fmt.Printf("Queue get: %v\n", item)
		}
	}()

	time.Sleep(1 * time.Second)
}

func NewQueue() *Queue {
	items := make(chan []Item, 1)
	empty := make(chan bool, 1)
	empty <- true
	return &Queue{items, empty}
}

func (q *Queue) Get() Item {
	// 起到获取锁的作用
	items := <-q.items

	item := items[0]
	items = items[1:]
	if len(items) == 0 {
		q.empty <- true
	} else {
		q.items <- items
	}
	return item
}

func (q *Queue) Put(item Item) {
	var items []Item
	select {
	case items = <-q.items:
	case <-q.empty:
	}

	items = append(items, item)
	q.items <- items
}
