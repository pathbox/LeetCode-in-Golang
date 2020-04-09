package LeetCode404

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	dfs(root)
	return sum
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { // 左子树叶子节点
		sum += root.Left.Val
	}

	dfs(root.Left)  // 继续递归左子树
	dfs(root.Right) // 继续递归右子树

}
