package LeetCode701

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Left, val)
		}
	}
	return root
}
