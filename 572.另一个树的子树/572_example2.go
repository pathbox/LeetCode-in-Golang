package LeetCode572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isSubtree(s.Left, t) || equal(s, t) || isSubtree(s.Right, t)
	/*
		if equal(s, t) {
			return true
		} else if isSubtree(s.Left, t) {
			return true
		} else if isSubtree(s.Right, t) {
			return true
		}
		return false
	*/
}

func equal(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return equal(s.Left, t.Left) && equal(s.Right, t.Right)
}
