package LeetCode0412

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs 前序
func solve(root *TreeNode, target int) int {
	ans := 0
	if root == nil {
		return ans
	}
	if root.Val == target {
		ans++
	}
	ans += solve(root.Left, target-root.Val)
	ans += solve(root.Right, target-root.Val)
	return ans
}

// dfs
func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	ans := 0
	ans += solve(root, sum)
	ans += dfs(root.Left, sum)
	ans += dfs(root.Right, sum)
	return ans
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return dfs(root, sum)
}

// 两个dfs
