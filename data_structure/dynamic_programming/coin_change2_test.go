package dynamic_programming

import "testing"

func TestCoinChange2(t *testing.T) {
	coins := []int{1, 2, 5}
	amt := 5
	got := CoinChange2(coins, amt)
	expected := 4
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
