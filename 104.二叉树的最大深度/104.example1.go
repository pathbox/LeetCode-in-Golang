package LeetCode104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := maxDepth(root.Left) // 每往下递归一次，高度其实就是加1
	rightHeight := maxDepth(root.Right)

	return max(leftHeight, rightHeight) + 1 // 返回当前左右子树最大的高度
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 每个节点访问一次，时间复杂度为 O(N)O(N)
