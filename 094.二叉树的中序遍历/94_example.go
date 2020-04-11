package LeetCode055

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	dfs(root, &nums)
	return nums
}

func dfs(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, nums)            // 先左子树
	*nums = append(*nums, root.Val) // 父节点
	dfs(root.Right, nums)           // 右子树
}
