package dynamic_programming

import "testing"

func TestMinDistance(t *testing.T) {
	// word1 = "horse", word2 = "ros"
	word1 := "horse"
	word2 := "ros"
	got := minDistance(word1, word2)
	expected := 3
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	// word1 = "intention", word2 = "execution"
	word1 = "intention"
	word2 = "execution"
	got = minDistance(word1, word2)
	expected = 5
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
