package Offer054

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树中序遍历是递增序列,然后取len(res) - k 位置的值，第一大就是最后一个数
func kthLargest(root *TreeNode, k int) int {
	res := make([]int, 0)
	dfs(root, &res)
	i := len(res) - k
	return res[i]
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	*res = append(*res, root.Val)
	dfs(root.Right, res)
}
