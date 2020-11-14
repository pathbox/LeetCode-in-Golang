package LeetCode783

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var pre int

// 中序遍历 + pre 指针值表示上一个节点，然后每次和root 当前节点做差
func minDiffInBST(root *TreeNode) int {
	res = 1<<63 - 1
	pre = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if pre != 0 {
		res = min(res, root.Val-pre)
	}
	pre = root.Val
	dfs(root.Right)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
