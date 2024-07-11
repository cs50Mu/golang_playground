package dynamic_programming

import "testing"

func TestBestTimeBuySell(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	got := BestTimeBuySell(prices)
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
	got = BestTimeBuySellTwoPointer(prices)
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	expected = 7
	got = btbss2(prices)
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
	prices = []int{7, 6, 4, 3, 1}
	expected = 0
	got = btbss2(prices)
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
