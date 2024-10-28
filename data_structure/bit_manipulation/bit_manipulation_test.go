package bitmanipulation

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBitSet(t *testing.T) {
	n := 1000
	round := 10000

	bs := NewBitSet(n)
	m := make(map[int]bool)

	fmt.Println("Prepare to do the test...")
	for i := 0; i < round; i++ {
		deside := rand.Float32()
		num := rand.Intn(n)
		if deside < 0.333 {
			bs.Add(num)
			m[num] = true
		} else if deside < 0.666 {
			bs.Remove(num)
			delete(m, num)
		} else {
			bs.Flip(num)
			if _, ok := m[num]; ok {
				delete(m, num)
			} else {
				m[num] = true
			}
		}
	}
	fmt.Println("test done, start checking result...")

	for i := 0; i < n; i++ {
		if bs.Contains(i) != m[i] {
			t.Errorf("Wrong result for %v", i)
		}
	}
}

func TestAdd(t *testing.T) {
	got := add(69, 69)
	want := 138
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestMinus(t *testing.T) {
	got := minus(69, 9)
	want := 60
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := multiply(-30, -5)
	want := 150
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestDiv(t *testing.T) {
	got := div(-60, -50)
	want := 1
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
