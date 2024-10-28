package stack_queue

import "testing"

func TestStack(t *testing.T) {
	assertEqual := func(want, got any) {
		if want != got {
			t.Errorf("want: %v, got: %v", want, got)
		}
	}

	stack := NewStack(2)

	var err error
	var want, got any
	_, err = stack.Pop()
	want = ErrStackEmpty
	assertEqual(want, err)

	err = stack.Push(1)
	want = nil
	assertEqual(want, err)

	size := stack.Size()
	want = 1
	assertEqual(want, size)

	stack.Push(2)
	got, _ = stack.Peek()
	want = 2
	assertEqual(want, got)

	stack.Pop()
	got, err = stack.Pop()
	want = nil
	assertEqual(want, err)
	want = 1
	assertEqual(want, got)

	_, err = stack.Pop()
	want = ErrStackEmpty
	assertEqual(want, err)
}
