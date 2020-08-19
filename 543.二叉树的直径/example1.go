package LeetCode543

import "math"

func diameterOfBinaryTree(root *TreeNode) int {
	m := 0
	max := &m //现有的最长直径
	if root == nil {
		return 0
	}
	depth(root, max)
	return *max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	a := depth(root.Left, max)
	b := depth(root.Right, max)
	*max = int(math.Max(float64(a+b), float64(*max))) //此节点的直径和现有的最长直径比较
	return int(math.Max(float64(a), float64(b))) + 1  //返回深度
}
