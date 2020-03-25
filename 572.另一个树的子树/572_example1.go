package LeetCode572

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	ss := preOrder(s, "")
	ts := preOrder(t, "")
	return strings.Contains(ss, ts)
}

func preOrder(tn *TreeNode, prefix string) string {
	if tn == nil {
		return prefix + "N"
	}
	return fmt.Sprintf("#%d%s%s", tn.Val, preOrder(tn.Left, "l"), preOrder(tn.Right, "r")) // 每次遍历的子树都打印出来
}
