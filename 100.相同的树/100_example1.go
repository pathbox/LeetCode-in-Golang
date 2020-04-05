package LeetCode100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { // 两者都为nil
		return true
	} else if p == nil || q == nil { // 只有一个位nil
		return false
	}

	if p.Val != q.Val { // 判断值是否相等
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) // 继续左右子树分别递归，并且要都相等
}
