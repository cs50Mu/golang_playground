package dynamic_programming

import "testing"

func TestCanPatition(t *testing.T) {
	a := []int{1, 5, 11, 5}

	got := canPatition(a)
	expected := true
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	a = []int{1, 2, 3, 5}
	got = canPatition(a)
	expected = false
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
