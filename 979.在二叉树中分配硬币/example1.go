package LeetCode979

// https://leetcode-cn.com/problems/distribute-coins-in-binary-tree/solution/tong-guo-xie-zi-jie-dian-di-gui-fu-jie-dian-ji-sua/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func distributeCoins(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	res += abs(l) + abs(r)
	return l + r + root.Val - 1 // 留下一个硬币，剩下的移走
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}
