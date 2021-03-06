package Offer032

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS 广度优先遍历
func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		length := len(stack)
		for length > 0 {
			node := stack[0]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			res = append(res, node.Val)
			length--
			stack = stack[1:]
		}
	}
	return res
}
