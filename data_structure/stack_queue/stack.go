package stack_queue

import (
	"errors"
)

// Stack stack implemented in arr
// s.arr[s.size] is the next slot to insert
type Stack struct {
	arr  []int
	size int
}

var (
	ErrStackFull  = errors.New("stack is full")
	ErrStackEmpty = errors.New("stack is empty")
)

func NewStack(n int) *Stack {
	return &Stack{
		arr: make([]int, n),
	}
}

func (s *Stack) Push(e int) error {
	if s.isFull() {
		return ErrStackFull
	}
	s.arr[s.size] = e
	s.size++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, ErrStackEmpty
	}
	s.size--
	return s.arr[s.size], nil
}

func (s *Stack) Peek() (int, error) {
	if s.isEmpty() {
		return 0, ErrStackEmpty
	}
	return s.arr[s.size-1], nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func (s *Stack) isFull() bool {
	return s.size == len(s.arr)
}

// MyQueue implement Queue using stacks
type MyQueue struct {
	in  *Stack
	out *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		in:  NewStack(100),
		out: NewStack(100),
	}
}

func (this *MyQueue) Push(x int) {
	this.in.Push(x)
	this.inToOut()
}

func (this *MyQueue) Pop() int {
	this.inToOut()
	res, _ := this.out.Pop()
	return res
}

// inToOut bring all the elements in `In` to `out`
// 1) do this when and only when out is empty
// 2) when doing this, must transfer *all* the elements
// of `In` into `Out`
func (this *MyQueue) inToOut() {
	if this.out.isEmpty() {
		for !this.in.isEmpty() {
			e, _ := this.in.Pop()
			this.out.Push(e)
		}
	}
}

func (this *MyQueue) Peek() int {
	this.inToOut()
	res, _ := this.out.Peek()
	return res
}

func (this *MyQueue) Empty() bool {
	return this.in.isEmpty() && this.out.isEmpty()
}

// MinStack ...
// https://leetcode.com/problems/min-stack/
// https://www.bilibili.com/video/BV15X4y177cM
// 使用两个栈，一个放正常数据，一个放最小值的
type MinStack struct {
	data *Stack
	min  *Stack
}

func Constructor() MinStack {
	return MinStack{
		data: NewStack(30000),
		min:  NewStack(30000),
	}
}

func (this *MinStack) Push(val int) {
	// data 一定要 Push
	this.data.Push(val)
	if this.data.isEmpty() {
		this.min.Push(val)
	} else {
		minTop, _ := this.min.Peek()
		if val >= minTop {
			this.min.Push(minTop)
		} else {
			this.min.Push(val)
		}
	}
}

func (this *MinStack) Pop() {
	this.min.Pop()
	this.data.Pop()
}

func (this *MinStack) Top() int {
	e, _ := this.data.Peek()
	return e
}

func (this *MinStack) GetMin() int {
	e, _ := this.min.Peek()
	return e
}
