package array

import "testing"

func TestMinSubArry(t *testing.T) {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	expected := 0
	got := minSubArrayLen(target, nums)

	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
