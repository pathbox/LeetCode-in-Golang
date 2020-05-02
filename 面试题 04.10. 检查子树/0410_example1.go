package LeetCode0410

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == t2 {
		return true
	}

	if t1 == nil && t2 != nil || t1 != nil && t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2) // t1从左右子树继续与t2比较
	}
	return checkSubTree(t1.Right, t2.Right) && checkSubTree(t1.Left, t2.Left) // 以当前节点为root节点，继续比较左右子树节点
}

// 树的节点是基本组成元素，但这个节点不是独立的，它有左右子树连接，可以从该节点出发继续走左子树或右子树
