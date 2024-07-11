package tree

import "fmt"

type BstNode struct {
	Val   int
	Left  *BstNode
	Right *BstNode
}

func BstInsertIter(root *BstNode, data int) {

}

func BstSearchIter(root *BstNode, data int) int {
	return -1
}

func BstInsert(root *BstNode, data int) *BstNode {
	if root == nil {
		return &BstNode{
			Val: data,
		}
	}

	if data == root.Val {
		//
	} else if data > root.Val {
		root.Right = BstInsert(root.Right, data)
	} else {
		root.Left = BstInsert(root.Left, data)
	}
	return root
}

func BstSearch(root *BstNode, target int) int {
	if root == nil {
		return -1
	}

	if target == root.Val {
		return target
	} else if target > root.Val {
		return BstSearch(root.Right, target)
	} else {
		return BstSearch(root.Left, target)
	}
}

func BstMin(root *BstNode) *BstNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return BstMin(root.Left)
	} else {
		return root
	}
}

func BstMax(root *BstNode) *BstNode {
	if root == nil {
		return nil
	}
	if root.Right != nil {
		return BstMax(root.Right)
	} else {
		return root
	}
}

// BstFloor 小于等于该键的最大键
func BstFloor(root *BstNode, data int) int {
	if root == nil {
		return -1
	}

	if data == root.Val {
		return data
	} else if data > root.Val {
		// 在右子树中找
		res := BstFloor(root.Right, data)
		// 若能找到，那么就是它了
		if res != -1 {
			return res
			// 若找不到，则就是 root 节点本身
		} else {
			return root.Val
		}
	} else {
		return BstFloor(root.Left, data)
	}
}

// BstCeiling 大于等于该键的最小键
func BstCeiling(root *BstNode, data int) int {
	if root == nil {
		return -1
	}

	if data == root.Val {
		return data
	} else if data < root.Val {
		// 在左子树中找
		res := BstCeiling(root.Left, data)
		// 若能找到，那么就是它了
		if res != -1 {
			return res
			// 若找不到，则就是 root 节点本身
		} else {
			return root.Val
		}
	} else {
		return BstCeiling(root.Right, data)
	}
}

func BstDeleteMin(root *BstNode) *BstNode {
	if root == nil {
		return nil
	}

	// 若没有左节点，则当前节点就是最小的了，需要删掉的就是它
	if root.Left == nil {
		return root.Right
		// 若有左节点，则递归调用，返回值需要赋值给 root.Left
	} else {
		root.Left = BstDeleteMin(root.Left)
	}
	return root
}

func BstDeleteMinIter(root *BstNode) {
	if root == nil {
		return
	}

	var parent *BstNode
	for root.Left != nil {
		parent = root
		root = root.Left
	}

	parent.Left = root.Right
}

func BstDeleteMax(root *BstNode) *BstNode {
	if root == nil {
		return nil
	}

	if root.Right == nil {
		return root.Left
	} else {
		root.Right = BstDeleteMax(root.Right)
	}
	return root
}

func BstTraversal(root *BstNode) {
	if root == nil {
		return
	}

	BstTraversal(root.Left)
	fmt.Println(root.Val)
	BstTraversal(root.Right)
}

func BstDelete(root *BstNode, target int) *BstNode {
	if root == nil {
		return nil
	}

	// 在右子树继续找
	if target > root.Val {
		root.Right = BstDelete(root.Right, target)
		// 在左子树继续找
	} else if target < root.Val {
		root.Left = BstDelete(root.Left, target)
		// 找到了，那就看怎么来删除它了
	} else {
		// 任意一个子树为空的情况，直接返回另一个子树即可
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 当两个子树都不为空的情况下
		tmp := root
		root = BstMin(tmp.Right)
		root.Left = tmp.Left
		root.Right = BstDeleteMin(tmp.Right)
	}
	return root
}
