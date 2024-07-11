package lockfree

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := New()

	for i := 1; i <= 9; i++ {
		s.Push(i)
	}

	for i := 1; i <= 9; i++ {
		v, err := s.Pop()
		if err != nil {
			t.Errorf("expected nil, got err: %v", err)
		}
		fmt.Printf("%v\n", v)
	}

	_, err := s.Pop()
	if err == nil {
		t.Errorf("expect error, got nil")
	}
}
