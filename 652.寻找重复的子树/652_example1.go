package LeetCode652

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	m := map[string]*TreeNode{}
	msame := map[*TreeNode]bool{}
	res := []*TreeNode{}
	getKey(root, m, msame)
	for node := range msame {
		res = append(res, node)
	}
	return res
}

func getKey(root *TreeNode, m map[string]*TreeNode, msame map[*TreeNode]bool) string {
	if root == nil {
		return ""
	}
	s := strconv.Itoa(root.Val)
	flag := false
	if root.Left != nil {
		s += "(" + getKey(root.Left, m, msame)
		flag = true
	}
	if root.Right != nil {
		if !flag {
			s += "("
		}
		s += "," + getKey(root.Right, m, msame)
		flag = true
	}
	if flag {
		s += ")"
	}
	if node, ok := m[s]; ok {
		msame[node] = true
	} else {
		m[s] = root
	}
	return s
}
