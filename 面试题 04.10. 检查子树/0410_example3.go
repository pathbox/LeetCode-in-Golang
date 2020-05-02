package LeetCode0410

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	ss := preOrder(t1, "")
	ts := preOrder(t2, "")
	return strings.Contains(ss, ts)
}

func preOrder(tn *TreeNode, prefix string) string {
	if tn == nil {
		return prefix + "N"
	}
	return fmt.Sprintf("#%d%s%s", tn.Val, preOrder(tn.Left, "l"), preOrder(tn.Right, "r"))
}
