package LeetCode124

var res int

func maxPathSum(root *TreeNode) (sum int) {
	if root == nil {
		return 0
	}
	res = -1 << 63
	dfs(root)
	return res
}

func dfs(root *TreeNode) (sum int) {
	if root == nil {
		return
	}
	lMax := max(0, dfs(root.Left))
	rMax := max(0, dfs(root.Right))
	res = max(res, root.Val+lMax+rMax)
	return root.Val + max(lMax, rMax)
}

func max(nums ...int) int {
	max := -1 << 63
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
