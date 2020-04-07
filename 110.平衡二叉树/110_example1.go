package LeetCode110

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1
平衡子树暗示了一个事实，每棵子树也是一个子问题。
*/

func GetMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(GetMax(root.Right)), float64(GetMax(root.Left)))) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(GetMax(root.Left))-float64(GetMax(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
