package LeetCode1302

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int
var sum int

func deepestLeavesSum(root *TreeNode) int {
	max, sum = 0, 0
	dfs(root, 0)
	return sum
}

// 遍历树的深度的变种
func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	// 判断深度 - 大于最大深度则重置结果值，等于则相加，小于不做任何操作
	if level > max {
		max = level
		sum = root.Val
	} else if max == level {
		sum += root.Val
	}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
