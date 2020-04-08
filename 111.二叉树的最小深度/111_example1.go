package LeetCode111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	if node.Right == nil && node.Left == nil { // 表示node是叶子节点
		if depth < min {
			min = depth
		}
		return
	}
	if node.Left != nil { // 不是叶子节点，从左子树开始往下递归，需要+1
		dfs(node.Left, depth+1)
	}

	if node.Right != nil { // 不是叶子节点，从右子树开始往下递归，需要+1
		dfs(node.Right, depth+1)
	}
}
