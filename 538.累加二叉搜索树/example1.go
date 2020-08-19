package LeetCode538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var add int

func convertBST(root *TreeNode) *TreeNode {
	add = 0
	dfs(root)
	return root
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Right)
		root.Val += add
		add = root.Val
		dfs(root.Left)
	}
}
