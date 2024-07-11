package dynamic_programming

import "testing"

func TestKnapsack(t *testing.T) {
	// 	Input: N = 3, W = 4, profit[] = {1, 2, 3}, weight[] = {4, 5, 1}
	// Output: 3
	capacity := 4
	profit := []int{1, 2, 3}
	weight := []int{4, 5, 1}

	got := Knapsack01(capacity, profit, weight, len(weight))
	expected := 3
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	capacity = 4
	profit = []int{4, 2, 3}
	weight = []int{2, 1, 3}

	got = Knapsack01(capacity, profit, weight, len(weight))
	expected = 6
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
