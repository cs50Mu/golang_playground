package dynamic_programming

import "testing"

func TestTargetSum(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	target := 3

	got := TargetSum(a, target)
	expected := 5
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
