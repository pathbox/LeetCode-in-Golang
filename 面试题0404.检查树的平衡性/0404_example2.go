package LeetCode0404

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 其实就是检查每个节点的左右子树是否平衡
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(high(root.Left)-high(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func high(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(high(root.Left), high(root.Right)) + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
