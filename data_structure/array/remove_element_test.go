package array

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	target := 3
	expected := 2
	got := removeElement(nums, target)
	fmt.Printf("nums: %+v\n", nums)
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	got = removeElementSlowFast(nums, target)
	fmt.Printf("nums: %+v\n", nums)
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
