package bst

import (
	"testing"
)

func TestIt(t *testing.T) {
	nums := []int{1, 2, 3, 5, 9}
	target := 10
	res := binarySearch(nums, target, 0, len(nums)-1)
	if res != -1 {
		t.Errorf("expected: %v, got: %v\n", -1, res)
	}

	res = binarySearchIter(nums, target)
	if res != -1 {
		t.Errorf("expected: %v, got: %v\n", -1, res)
	}

	target = 1
	res = binarySearch(nums, target, 0, len(nums)-1)
	expected := 0
	if res != expected {
		t.Errorf("expected: %v, got: %v\n", expected, res)
	}

	res = binarySearchIter(nums, target)
	if res != expected {
		t.Errorf("expected: %v, got: %v\n", expected, res)
	}
}
