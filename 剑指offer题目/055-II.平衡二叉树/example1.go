package Offer056

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return root == nil || isBalanced(root.Left) && math.Abs(height(root.Left)-height(root.Right)) < 2 && isBalanced(root.Right)
}

//计算当前节点的最大深度
func height(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(height(node.Left), height(node.Right)) + 1
}
