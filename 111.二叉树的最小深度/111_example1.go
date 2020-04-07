package LeetCode111

var min int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min = 1<<63 - 1 // 初始化一个较大的值
	dfs(root, 1)
	return min
}

func dfs(node *TreeNode, depth int) {
	if node.Right == nil && node.Left == nil {
		if depth < min {
			min = depth
		}
		return
	}
	if node.Left != nil {
		dfs(node.Left, depth+1)
	}

	if node.Right != nil {
		dfs(node.Right, depth+1)
	}
}
