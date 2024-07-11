package dynamic_programming

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amt := 11
	got := CoinChange(coins, amt)
	expected := 3
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	coins = []int{3, 7, 405, 436}
	amt = 8839
	got = CoinChange(coins, amt)
	expected = 25
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
