package LeetCode0404

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return isBalancedHelper(root) != -1
}

func isBalancedHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := isBalancedHelper(root.Left)
	r := isBalancedHelper(root.Right)
	if l != -1 && r != -1 && abs(l-r) <= 1 {
		return max(l, r) + 1 // 求树的高度就是遍历左子树和右子树，求左右子树最大的高度+1
	}
	return -1
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
