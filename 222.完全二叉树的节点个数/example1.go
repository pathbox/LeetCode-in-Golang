package LeetCode222

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// 时间复杂度： O(N)。
// 空间复杂度： O(d) = O(logN)，其中 d 指的是树的的高度，运行过程中堆栈所使用的空间
