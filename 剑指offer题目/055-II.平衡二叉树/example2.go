package Offer056

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return dfsIsBalanced(root) != -1
}

func dfsIsBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfsIsBalanced(root.Left)
	r := dfsIsBalanced(root.Right)
	if l != -1 && r != -1 && int(math.Abs(float64(l-r))) <= 1 {
		return int(math.Max(float64(l), float64(r))) + 1
	}
	return -1
}

// 树的递归其实是自顶向下遍历，自底向上递归。也就是，最开始出栈进行逻辑操作的，是最后一个树的节点，然后往上到其父节点，然后继续往上其祖先节点
