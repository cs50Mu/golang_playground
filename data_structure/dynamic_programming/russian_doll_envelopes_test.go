package dynamic_programming

import "testing"

func TestMaxEnvlopes(t *testing.T) {
	// [[5,4],[6,4],[6,7],[2,3]]
	evps := [][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}

	got := maxEnvlopes(evps)
	expected := 3
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	evps = [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}

	got = maxEnvlopes(evps)
	expected = 1
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
