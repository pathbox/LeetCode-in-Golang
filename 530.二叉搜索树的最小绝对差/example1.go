package LeetCode530

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre int
var min int

// 中序遍历是递增数列顺序，相邻相减可以得到最小的差值
func getMinimumDifference(root *TreeNode) int {
	pre = -1
	min = math.MaxInt64
	midOrder(root)
	return min
}

// 中序遍历：左 root 右
func midOrder(root *TreeNode) {
	if root == nil {
		return
	}
	midOrder(root.Left)
	if pre != -1 {
		if root.Val-pre < min {
			min = root.Val - pre
		}
	}
	pre = root.Val
	midOrder(root.Right)
}
