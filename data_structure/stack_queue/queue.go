package stack_queue

import (
	"errors"
	"fmt"
)

type listNode struct {
	Val  int
	Next *listNode
}

// ListQueue queue implemented in linkedList
// when Head and Tail both
// are nil, the queue is empty
type ListQueue struct {
	// Dequeue operates on Head
	Head *listNode
	// Enqueue operates on Tail
	Tail *listNode
}

func newNode(val int) *listNode {
	return &listNode{
		Val: val,
	}
}

func NewListQueue() *ListQueue {
	return &ListQueue{}
}

func (q *ListQueue) Enqueue(e int) {
	if q.Tail == nil {
		q.Head = newNode(e)
		q.Tail = q.Head
	} else {
		q.Tail.Next = newNode(e)
		q.Tail = q.Tail.Next
	}
}

func (q *ListQueue) Dequeue() int {
	if q.Head != nil {
		headVal := q.Head.Val
		q.Head = q.Head.Next
		if q.Head == nil {
			// queue is empty
			// set tail to nil also
			q.Tail = nil
		}
		return headVal
		// empty
	} else {
		return -1
	}
}

func (q *ListQueue) Print() {
	var res []int
	curr := q.Head
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}
	fmt.Printf("%+v\n", res)
}

// ArrayQueue circular queue implemented in array
// ref: https://www.bilibili.com/video/BV18X4y177vA/?share_source=copy_web&vd_source=8f331b71355e4192b786cc4504fa3adc&t=494
// left and right are indexes
// 左开右闭，right指向下一个空闲的空间
// [left .... right)
type ArrayQueue struct {
	arr   []int
	left  int
	right int
	size  int // the current size of the queue
	cap   int // the capacity of the queue
}

var (
	ErrQueueIsFull  = errors.New("queue is full")
	ErrQueueIsEmpty = errors.New("queue is empty")
)

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		arr: make([]int, n),
		cap: n,
	}
}

// Enqueue ...
// operates on q.right
func (q *ArrayQueue) Enqueue(e int) error {
	if q.isFull() {
		return ErrQueueIsFull
	}
	q.arr[q.right] = e
	// wraps to 0 when reaches the arr bound
	if q.right == q.cap-1 {
		q.right = 0
	} else {
		q.right++
	}
	q.size++
	return nil
}

// Dequeue ...
// operates on q.left
func (q *ArrayQueue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, ErrQueueIsEmpty
	}
	e := q.arr[q.left]
	// wraps to 0 when reaches the arr bound
	if q.left == q.cap-1 {
		q.left = 0
	} else {
		q.left++
	}
	q.size--

	return e, nil
}

func (q *ArrayQueue) isEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue) isFull() bool {
	return q.size == q.cap
}

func (q *ArrayQueue) Peek() (int, error) {
	if q.isEmpty() {
		return 0, ErrQueueIsEmpty
	}

	return q.Head()
}

func (q *ArrayQueue) Head() (int, error) {
	if q.isEmpty() {
		return 0, ErrQueueIsEmpty
	}
	return q.arr[q.left], nil
}

func (q *ArrayQueue) Tail() (int, error) {
	if q.isEmpty() {
		return 0, ErrQueueIsEmpty
	}
	var prev int
	if q.right == 0 {
		prev = q.cap - 1
	} else {
		prev = q.right - 1
	}
	return q.arr[prev], nil
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) Print() {
	fmt.Printf("%+v\n", q.arr[q.left:q.right])
}

// MyStack implement stack using queue
// ref: https://www.bilibili.com/video/BV1E14y1B7j4/
type MyStack struct {
	Q *ArrayQueue
}

func Constructor() MyStack {
	return MyStack{
		Q: NewArrayQueue(100),
	}
}

// Push ...
// 先看内部队列里还有多少个元素，假定为 n 个
// 然后再 Push 进队列
// 最后，循环执行 n 次：Dequeue / Enqueue
func (this *MyStack) Push(x int) {
	n := this.Q.Size()
	this.Q.Enqueue(x)
	for ; n > 0; n-- {
		e, _ := this.Q.Dequeue()
		this.Q.Enqueue(e)
	}
}

func (this *MyStack) Pop() int {
	e, _ := this.Q.Dequeue()
	return e
}

// Top 栈顶即队列顶
func (this *MyStack) Top() int {
	e, _ := this.Q.Peek()
	return e
}

func (this *MyStack) Empty() bool {
	return this.Q.isEmpty()
}

// DoubleEndQueue 双端队列
// 使用数组实现
// https://leetcode.com/problems/design-circular-deque/description/
// ref: https://www.bilibili.com/video/BV1PM4y1p7N5/
type DoubleEndQueue struct {
	arr  []int
	l, r int // left and right boundary of the queue
	size int // the size of the queue
	cap  int // the capicity of the queue
}

func NewDoubleEndQueue(n int) *DoubleEndQueue {
	return &DoubleEndQueue{
		arr: make([]int, n),
		cap: n,
	}
}

func (deq *DoubleEndQueue) PushHead(e int) error {
	if deq.isFull() {
		return ErrQueueIsFull
	}
	if deq.isEmpty() {
		// need to reset l and r to 0
		deq.l = 0
		deq.r = deq.l
		deq.arr[deq.l] = e
	} else {
		if deq.l == 0 {
			deq.l = deq.cap - 1
		} else {
			deq.l -= 1
		}
		deq.arr[deq.l] = e
	}
	deq.size++
	return nil
}

func (deq *DoubleEndQueue) PopHead() (int, error) {
	if deq.isEmpty() {
		return 0, ErrQueueIsEmpty
	}
	res := deq.arr[deq.l]
	if deq.l == deq.cap-1 {
		deq.l = 0
	} else {
		deq.l += 1
	}
	deq.size--
	return res, nil
}

func (deq *DoubleEndQueue) PushTail(e int) error {
	if deq.isFull() {
		return ErrQueueIsFull
	}
	if deq.isEmpty() {
		// need to reset l and r to 0
		deq.l = 0
		deq.r = deq.l
		deq.arr[deq.l] = e
	} else {
		if deq.r == deq.cap-1 {
			deq.r = 0
		} else {
			deq.r += 1
		}
		deq.arr[deq.r] = e
	}
	deq.size++
	return nil

}

func (deq *DoubleEndQueue) PopTail() (int, error) {
	if deq.isEmpty() {
		return 0, ErrQueueIsEmpty
	}
	res := deq.arr[deq.r]
	if deq.r == 0 {
		deq.r = deq.cap - 1
	} else {
		deq.r -= 1
	}
	deq.size--
	return res, nil
}

// Head returns the element for the Head index
func (deq *DoubleEndQueue) Head() (int, error) {
	if deq.isEmpty() {
		return 0, ErrQueueIsEmpty
	}
	return deq.arr[deq.l], nil
}

// Tail returns the element for the Tail index
func (deq *DoubleEndQueue) Tail() (int, error) {
	if deq.isEmpty() {
		return 0, ErrQueueIsEmpty
	}
	return deq.arr[deq.r], nil
}

func (deq *DoubleEndQueue) isEmpty() bool {
	return deq.size == 0
}

func (deq *DoubleEndQueue) isFull() bool {
	return deq.size == deq.cap
}
