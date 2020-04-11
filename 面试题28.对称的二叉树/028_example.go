package LeetCode700

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSame(root.Left, root.Right)
}

func isSame(L, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}
	if L == nil && R != nil {
		return false
	}
	if L != nil && R == nil {
		return false
	}
	if L.Val != R.Val {
		return false
	}
	return isSame(L.Left, R.Right) && isSame(L.Right, R.Left)
}
