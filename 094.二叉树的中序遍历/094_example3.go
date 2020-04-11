package LeetCode094

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		// push
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		pre := len(stack) - 1
		res = append(res, stack[pre].Val)
		root = stack[pre].Right
		stack = stack[:pre]
	}
	return res
}
