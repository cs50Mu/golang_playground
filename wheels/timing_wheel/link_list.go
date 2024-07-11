package timing_wheel

type Node[K, V any] struct {
	Key  K
	Val  V
	Prev *Node[K, V]
	Next *Node[K, V]
}

type DoublyLinkedList[K, V any] struct {
	Head *Node[K, V]
	Tail *Node[K, V]
	Len  int64
}

func NewDoublyLinkedList[K, V any]() *DoublyLinkedList[K, V] {
	head := newDummyNode[K, V]()
	tail := newDummyNode[K, V]()
	head.Next = tail
	tail.Prev = head
	return &DoublyLinkedList[K, V]{
		Head: head,
		Tail: tail,
	}
}

func newDummyNode[K, V any]() *Node[K, V] {
	return &Node[K, V]{}
}

func newNode[K, V any](key K, val V) *Node[K, V] {
	return &Node[K, V]{
		Key: key,
		Val: val,
	}
}

func (dll *DoublyLinkedList[K, V]) PushBack(key K, val V) {
	node := newNode(key, val)
	prev := dll.Tail.Prev
	prev.Next = node
	node.Next = dll.Tail
	dll.Tail.Prev = node
	node.Prev = prev
	dll.Len++
}

func (dll *DoublyLinkedList[K, V]) Remove(node *Node[K, V]) {
	prev := node.Prev
	prev.Next = node.Next
	node.Next.Prev = prev
	// node.Prev = nil
	// node.Next = nil
	dll.Len--
}

func (dll *DoublyLinkedList[K, V]) Keys() []K {
	curr := dll.Head.Next
	var res []K
	for curr != dll.Tail {
		res = append(res, curr.Key)
		curr = curr.Next
	}

	return res
}
