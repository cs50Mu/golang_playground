package tree

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

type node struct {
	Val   int
	Left  *node
	Right *node
}

func preOrderIter(root *node) {
	s := stack.New()

	s.Push(root)
	for s.Len() > 0 {
		nn := s.Pop()
		node, _ := nn.(*node)
		fmt.Println(node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
}

func preOrder(root *node) {
	if root == nil {
		return
	}

	fmt.Printf("%v\n", root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *node) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	fmt.Printf("%v\n", root.Val)
	inOrder(root.Right)
}

func levelOrder(root *node) {
	fmt.Println("\nlevel order")
	if root == nil {
		return
	}
	var queue []*node
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%v\n", node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
