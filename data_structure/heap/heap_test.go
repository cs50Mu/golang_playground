package heap

import "testing"

func TestMaxPQ(t *testing.T) {
	mpq := NewMaxPQ(100)

	mpq.Insert(5)
	mpq.Insert(10)
	mpq.Insert(3)

	got := mpq.DelMax()
	expected := 10
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
