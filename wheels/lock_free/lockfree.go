package lockfree

import (
	"errors"
	"sync"
	"sync/atomic"
)

type Node struct {
	Data int
	Next *Node
}

func newNode(v int) *Node {
	return &Node{
		Data: v,
	}
}

func New() *LStack {
	return &LStack{}
}

var ErrEmptyStack = errors.New("stack is empty")

// LStack is a stack with lock
type LStack struct {
	mu  sync.Mutex
	Top *Node
}

func (s *LStack) Push(v int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	node := newNode(v)
	node.Next = s.Top
	s.Top = node
}

func (s *LStack) Pop() (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Top == nil {
		return 0, ErrEmptyStack
	}

	top := s.Top
	s.Top = top.Next
	// help gc
	top.Next = nil
	return top.Data, nil
}

type LFStack struct {
	Top atomic.Pointer[Node]
}

func (s *LFStack) Push(v int) {

}

func (s *LFStack) Pop() (int, error) {
	return 0, nil
}
