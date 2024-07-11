package string

import "testing"

func TestReplaceWhitespace(t *testing.T) {
	s := "We are happy."
	got := replaceWhitespace(s)
	expected := "We%20are%20happy."
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
