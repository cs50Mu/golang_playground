package stack_queue

import "testing"

func TestListQueue(t *testing.T) {
	lq := NewListQueue()

	lq.Enqueue(1)
	lq.Enqueue(2)
	lq.Enqueue(3)
	lq.Print()
	lq.Dequeue()
	lq.Dequeue()
	lq.Print()
	lq.Enqueue(4)
	lq.Print()
	lq.Dequeue()
	lq.Dequeue()
	lq.Print()
	lq.Enqueue(5)
	lq.Enqueue(6)
	lq.Print()
}

func TestArrayQueue(t *testing.T) {
	assertEqual := func(want, got any) {
		if want != got {
			t.Errorf("want: %v, got: %v", want, got)
		}
	}
	aq := NewArrayQueue(50)

	aq.Enqueue(1)
	aq.Enqueue(2)
	aq.Enqueue(3)
	aq.Print()
	aq.Dequeue()
	aq.Dequeue()
	aq.Print()
	aq.Enqueue(4)
	aq.Print()
	aq.Dequeue()
	aq.Dequeue()
	aq.Print()
	aq.Enqueue(5)
	aq.Enqueue(6)
	aq.Print()

	var err error
	var want, got any

	got, err = aq.Head()
	assertEqual(nil, err)
	want = 5
	assertEqual(want, got)
	aq.Dequeue()
	aq.Dequeue()
	_, err = aq.Tail()
	want = ErrQueueIsEmpty
	assertEqual(want, err)
}
