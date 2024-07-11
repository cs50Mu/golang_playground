package delay_queue

import (
	"testing"
)

func TestMinPQ(t *testing.T) {
	maxPQ := NewPQ[int](10, PQTypeMax)

	maxPQ.Insert(8)
	maxPQ.Insert(5)
	maxPQ.Insert(20)
	maxPQ.Insert(2)

	res := maxPQ.DelTop()

	expected := 20
	if res != expected {
		t.Errorf("got: %v, expected: %v", res, expected)
	}

	minPQ := NewPQ[int](10, PQTypeMin)

	minPQ.Insert(8)
	minPQ.Insert(5)
	minPQ.Insert(20)
	minPQ.Insert(2)

	res = minPQ.DelTop()

	expected = 2
	if res != expected {
		t.Errorf("got: %v, expected: %v", res, expected)
	}
}
