package LeetCode543

var max int

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	if root == nil {
		return 0
	}

	depth(root)
	return max
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := depth(root.Left)
	r := depth(root.Right)
	max = maxF(l+r, max)
	return maxF(l, r) + 1
}

func maxF(a, b int) int {
	if a > b {
		return a
	}
	return b
}
