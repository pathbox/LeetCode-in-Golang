package LeetCode701

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	r := root
	for {
		if r.Val > val {
			if r.Left == nil {
				r.Left = &TreeNode{Val: val}
				break
			} else {
				r = r.Left
			}
		} else {
			if r.Right == nil {
				r.Right = &TreeNode{val, nil, nil}
				break
			} else {
				r = r.Right
			}
		}
	}
	return root
}
