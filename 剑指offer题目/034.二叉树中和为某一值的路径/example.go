package Offer034

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	dfs(root, sum, []int{})
	return res
}

func dfs(root *TreeNode, sum int, arr []int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		res = append(res, tmp)
	}
	dfs(root.Left, sum-root.Val, arr)
	dfs(root.Right, sum-root.Val, arr)
	arr = arr[:len(arr)-1]
}
