package linked_list

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func NewDLL() *DoublyLinkedList {
	ddl := &DoublyLinkedList{}
	ddl.Head = &Node{}
	ddl.Tail = &Node{}

	ddl.Head.Next = ddl.Tail
	ddl.Tail.Prev = ddl.Head

	return ddl
}

func (dl *DoublyLinkedList) RemoveAt(node *Node) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}

func (dl *DoublyLinkedList) PopFront() *Node {
	node := dl.Head.Next
	dl.RemoveAt(node)
	return node
}

func (dl *DoublyLinkedList) RemoveLast() {
	last := dl.Tail.Prev
	dl.RemoveAt(last)
}

func (dl *DoublyLinkedList) PushBack(node *Node) {
	prev := dl.Tail.Prev
	prev.Next = node
	node.Prev = prev

	node.Next = dl.Tail
	dl.Tail.Prev = node
}

func (dl *DoublyLinkedList) AppendFront(node *Node) {
	next := dl.Head.Next

	dl.Head.Next = node
	node.Prev = dl.Head

	node.Next = next
	next.Prev = node
}

func (dl *DoublyLinkedList) Print() {
	fmt.Println("ddl now:")
	curr := dl.Head.Next
	for curr != dl.Tail {
		fmt.Println(curr.Val)
		curr = curr.Next
	}

	fmt.Println()
}
