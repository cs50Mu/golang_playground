package graph

import (
	"fmt"
	"testing"
)

func TestDigraph(t *testing.T) {
	tmp := initGraph(t, NewDigraph, "tinyDG.txt")
	g := tmp.(*Digraph)
	fmt.Printf("graph: %+v\n", *g)
	dd := NewDirectedDFS(g, 0)

	got := dd.Marked(11)
	expected := false
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	got = dd.Marked(3)
	expected = true
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	dc := NewDirectedCycle(g)
	got = dc.HasCycle()
	expected = true
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	fmt.Printf("cycles: %+v\n", dc.Cycle())

	// TODO: test case for topological sort
}
