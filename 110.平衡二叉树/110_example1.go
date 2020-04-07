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

1.从顶到底：
该方法是从顶到底进行判断，先判断当前节点是否平衡（分别求出左子树和右子树的高度，然后比较是否相差大于1），然后进行递归判断左子树和右子树是否平衡。
该方法的时间复杂度为O(nlogn)，缺陷是每个节点的高度会重复计算

*/

func isBalanced(node *TreeNode) bool {
	return node == nil || isBalanced(node.Left) &&
		math.Abs(height(node.Left)-height(node.Right)) < 2 && //两边最大深度差  大于2
		isBalanced(node.Right)
}

//计算节点最大深度
func height(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(height(node.Left), height(node.Right)) + 1
}
