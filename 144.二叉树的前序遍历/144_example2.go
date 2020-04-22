package LeetCode144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// stack迭代法
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root} // 当前节点入栈
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		res = append(res, p.Val)
		if p.Right != nil { // 先让右子树进栈，左子树后进栈，遍历的时候，就会先取左子树
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}

	return res
}
