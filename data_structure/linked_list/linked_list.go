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

func (dl *DoublyLinkedList) RemoveLast() *Node {
	last := dl.Tail.Prev
	dl.RemoveAt(last)
	return last
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

// dll 双向链表实现
type dll struct {
	head *Node
	tail *Node
}

func NewDll2() *dll {
	return &dll{
		head: nil,
		tail: nil,
	}
}

func newNode(e int) *Node {
	return &Node{
		Val: e,
	}
}

func (l *dll) PushHead(e int) {
	node := newNode(e)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.Next = l.head
		l.head.Prev = node
		l.head = l.head.Prev
	}
}

func (l *dll) PopHead() int {
	if l.head == nil {
		panic("queue is empty")
	}
	res := l.head.Val
	next := l.head.Next
	if next == nil {
		// reset head && tail
		l.head = nil
		l.tail = nil
	} else {
		// remove links
		l.head.Next = nil
		next.Prev = nil
		l.head = next
	}
	return res
}

func (l *dll) PushTail(e int) {
	node := newNode(e)
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		node.Prev = l.tail
		l.tail = l.tail.Next
	}
}

func (l *dll) PopTail() int {
	if l.tail == nil {
		panic("queue is empty")
	}
	res := l.tail.Val
	prev := l.tail.Prev
	if prev == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.tail.Prev = nil
		prev.Next = nil
		l.tail = prev
	}
	return res
}

func (l *dll) Print() {
	fmt.Print("nodes: ")
	for curr := l.head; curr != nil; curr = curr.Next {
		fmt.Printf("%v, ", curr.Val)
	}
	fmt.Println()
}

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

func newListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currA := headA
	currB := headB

	var sumA, sumB int
	var lastA, lastB *ListNode
	for currA != nil {
		sumA += 1
		lastA = currA
		currA = currA.Next
	}
	for currB != nil {
		sumB += 1
		lastB = currB
		currB = currB.Next
	}

	if lastA != lastB {
		return nil
	}

	var long, short *ListNode
	var preRun int
	if sumA > sumB {
		preRun = sumA - sumB
		long = headA
		short = headB
	} else {
		preRun = sumB - sumA
		long = headB
		short = headA
	}

	for i := 0; i < preRun; i++ {
		long = long.Next
	}

	for long != short {
		long = long.Next
		short = short.Next
	}

	return long
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	var prev *ListNode
	var i int
	for i = k; curr != nil && i > 0; i-- {
		prev = curr
		curr = curr.Next
	}
	if curr == nil {
		if i > 0 {
			return head
		} else {
			return reverse(head)
		}
	}

	prev.Next = nil
	newHead := reverse(head)
	lastTail := head
	fmt.Printf("lastTail: %v\n", lastTail.Val)

	for curr != nil {
		xxx(&curr, k, &lastTail)
		// oldHead := curr // 先记一下老的头节点，反转之后它会是尾
		// fmt.Printf("oldHead: %v\n", oldHead.Val)
		// var i int
		// // 跳过下 k 个元素
		// for i = k; curr != nil && i > 0; i-- {
		// 	prev = curr
		// 	curr = curr.Next
		// }
		// // 若当前组不够 k 个了，就直接退出
		// // 但要跟之前的尾连一下先
		// if i > 0 && curr == nil {
		// 	lastTail.Next = oldHead
		// 	break
		// }

		// prev.Next = nil // prev 指向的是 k group 的最后一个元素
		// // 置空是为了让 reverse 可以正常工作
		// h := reverse(oldHead)
		// // 把新反转的 group 跟之前的尾接上
		// lastTail.Next = h
		// // 更新 lastTail
		// // 它应该是当前 k group 在没有反转之前的头
		// lastTail = oldHead
	}

	return newHead
}

// 循环里的代码可以提取出来，但感觉也不太能跟循环之外的代码复用
// 因为第一个 k group 需要特殊处理（因为head），最关键的是，
// 我下次写肯定还是直接写出没有提取出函数的版本
func xxx(curr **ListNode, k int, lastTail **ListNode) {
	oldHead := *curr // 先记一下老的头节点，反转之后它会是尾
	fmt.Printf("oldHead: %v\n", oldHead.Val)
	var i int
	// 跳过下 k 个元素
	var prev *ListNode
	for i = k; *curr != nil && i > 0; i-- {
		prev = *curr
		*curr = (*curr).Next
	}
	// 若当前组不够 k 个了，就直接退出
	// 但要跟之前的尾连一下先
	if i > 0 && *curr == nil {
		(*lastTail).Next = oldHead
		return
	}

	prev.Next = nil // prev 指向的是 k group 的最后一个元素
	// 置空是为了让 reverse 可以正常工作
	h := reverse(oldHead)
	// 把新反转的 group 跟之前的尾接上
	(*lastTail).Next = h
	// 更新 lastTail
	// 它应该是当前 k group 在没有反转之前的头
	*lastTail = oldHead
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func copyRandomList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	// make 1->2->3 becomes
	// 1->1'->2->2'->3->3'
	for curr != nil {
		newNode := &ListNode{
			Val:    curr.Val,
			Random: curr.Random,
		}
		next := curr.Next
		curr.Next = newNode
		newNode.Next = next
		curr = next
	}
	// copy random links
	curr = head
	for curr != nil {
		newNode := curr.Next
		oldRand := curr.Random
		if oldRand != nil {
			newNode.Random = oldRand.Next
		}
		curr = curr.Next.Next
	}

	printListNode(head)

	// seperate the two lists
	curr = head

	ans := head.Next // 新链表的头
	for curr != nil {
		next := curr.Next.Next
		newNode := curr.Next
		curr.Next = curr.Next.Next
		if next == nil {
			newNode.Next = nil
		} else {
			newNode.Next = newNode.Next.Next
		}
		curr = next
	}

	return ans
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// find the mid using slow fast pointers
	slow := head
	fast := slow
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	// only one element
	if fast == slow {
		return true
	}

	// reverse the right part of the list
	h := reverse(slow)

	// check palindrome
	currLeft := head
	currRight := h
	ans := true
	for currLeft != nil && currRight != nil {
		if currLeft.Val != currRight.Val {
			ans = false
			break
		}
		currLeft = currLeft.Next
		currRight = currRight.Next
	}

	// restore the original list
	restored := reverse(h)
	prev.Next = restored

	return ans
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := slow
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	// there is no circle in this list
	if fast == nil || fast.Next == nil {
		return nil
	}

	// reset fast to head
	// and slow and fast both
	// move one step every time
	fast = head
	for {
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
}
