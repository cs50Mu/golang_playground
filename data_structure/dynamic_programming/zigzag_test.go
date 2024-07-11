package dynamic_programming

import "testing"

func TestZigZag(t *testing.T) {
	seq := []int{1, 7, 4, 9, 2, 5}

	got := zigZag(seq)
	expected := 6
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}

	seq = []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	got = zigZag(seq)
	expected = 7
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}

	seq = []int{55}
	got = zigZag(seq)
	expected = 1
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}

	seq = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got = zigZag(seq)
	expected = 2
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}

	seq = []int{70, 55, 13, 2, 99, 2, 80, 80, 80, 80, 100, 19, 7, 5, 5, 5, 1000, 32, 32}
	got = zigZag(seq)
	expected = 8
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
