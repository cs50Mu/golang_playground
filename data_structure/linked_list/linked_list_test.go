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

// TestDLL2 ...
// dll2 implement doubly linked list without dummy nodes
func TestDLL2(t *testing.T) {
	dll := NewDll2()

	dll.PushHead(1)
	dll.PushHead(2)
	dll.Print()
	// dll.PopHead()
	dll.PushTail(3)
	dll.PushTail(4)
	dll.Print()
	dll.PopTail()
	dll.Print()
	dll.PopHead()
	dll.Print()
	dll.PopTail()
	dll.PopTail()
	dll.Print()
}

func TestReverse(t *testing.T) {
	head := newListNode(1)
	curr := head
	curr.Next = newListNode(2)
	curr = curr.Next
	curr.Next = newListNode(3)
	curr = curr.Next
	curr.Next = newListNode(4)
	curr = curr.Next
	curr.Next = newListNode(5)
	printListNode(head)

	head = reverseKGroup(head, 2)
	printListNode(head)
}

func TestCopyRandList(t *testing.T) {
	head := newListNode(1)
	curr := head
	curr.Next = newListNode(2)
	curr = curr.Next
	curr.Next = newListNode(3)
	printListNode(head)
	h := copyRandomList(head)
	printListNode(h)
}
