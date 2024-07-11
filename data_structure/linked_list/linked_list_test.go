package linked_list

import "testing"

func TestLinkedList(t *testing.T) {
	dll := NewDLL()

	node1 := newNode(1)
	node2 := newNode(2)
	node3 := newNode(3)

	dll.AppendFront(node1)
	dll.AppendFront(node2)
	dll.AppendFront(node3)

	dll.Print()

	dll.RemoveLast()

	dll.Print()

	// dll.RemoveAt(node2)

	// dll.Print()

	// dll.RemoveAt(node3)

	// dll.Print()

	// dll.RemoveAt(node1)

	// dll.Print()
}

func newNode(val int) *Node {
	return &Node{
		Val: val,
	}
}
