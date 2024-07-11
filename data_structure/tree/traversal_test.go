package tree

import (
	"fmt"
	"testing"
)

func TestTraversal(t *testing.T) {
	rNode := node{
		Val: 0,
	}

	leftNode := node{
		Val: 1,
	}
	rightNode := node{
		Val: 2,
	}

	rNode.Left = &leftNode
	rNode.Right = &rightNode

	preOrder(&rNode)
	fmt.Println()
	preOrderIter(&rNode)

	fmt.Println()

	inOrder(&rNode)

	levelOrder(&rNode)
}
