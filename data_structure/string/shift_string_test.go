package string

import "testing"

func TestShiftString(t *testing.T) {
	// s = "abcdefg", k = 2
	s := "abcdefg"
	k := 2
	got := shiftString(s, k)
	expected := "cdefgab"
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
