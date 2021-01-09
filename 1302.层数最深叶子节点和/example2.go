package LeetCode1302

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	var count int
	for len(stack) > 0 {
		count = 0
		length := len(stack)
		for i := 0; i < length; i++ {
			node := stack[i]
			count += node.Val
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		stack = stack[length:]
	}
	return count
}
