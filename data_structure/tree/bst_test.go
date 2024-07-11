package tree

import (
	"fmt"
	"testing"
)

func TestBst(t *testing.T) {
	root := &BstNode{
		Val: 5,
	}

	BstInsert(root, 1)
	BstInsert(root, 7)
	BstInsert(root, 2)
	BstInsert(root, 3)

	res := BstSearch(root, 7)
	expected := 7
	if res != expected {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	res = BstSearch(root, 17)
	expected = -1
	if res != expected {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	resN := BstMin(root)
	expected = 1
	if resN.Val != expected {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	resN = BstMax(root)
	expected = 7
	if resN.Val != expected {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	res = BstFloor(root, 4)
	expected = 3
	if res != expected {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	res = BstFloor(root, 6)
	expected = 5
	if res != expected {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	res = BstCeiling(root, 6)
	expected = 7
	if res != expected {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	res = BstCeiling(root, 4)
	expected = 5
	if res != expected {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

}

func TestBstDelete(t *testing.T) {
	root := &BstNode{
		Val: 5,
	}

	root = BstInsert(root, 1)
	root = BstInsert(root, 7)
	root = BstInsert(root, 2)
	root = BstInsert(root, 3)

	BstTraversal(root)

	root = BstDeleteMin(root)

	fmt.Println()
	BstTraversal(root)

	BstDeleteMinIter(root)

	fmt.Println()
	BstTraversal(root)

	root = BstDeleteMax(root)
	fmt.Println()
	BstTraversal(root)

	root = BstDelete(root, 5)
	fmt.Println()
	BstTraversal(root)
}
