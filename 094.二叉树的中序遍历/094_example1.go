package LeetCode055

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
时间复杂度：O(n)O(n)。递归函数 T(n) = 2 \cdot T(n/2)+1T(n)=2⋅T(n/2)+1。
空间复杂度：最坏情况下需要空间O(n)O(n)，平均情况为O(log n)O(logn)
*/
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
