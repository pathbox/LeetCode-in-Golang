package LeetCode652

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res, m := make([]*TreeNode, 0), make(map[string]int)
	helper(root, &res, m)
	return res
}

func helper(root *TreeNode, res *[]*TreeNode, m map[string]int) string {
	if root == nil {
		return ""
	}

	route := fmt.Sprintf("%s#%s#%s",
		root.Val, helper(root.Left, res, m), helper(root.Right, res, m))

	switch {
	case m[route] == 0:
		m[route]++
	case m[route] == 1:
		m[route]++
		*res = append(*res, root)
	}
	return route
}
