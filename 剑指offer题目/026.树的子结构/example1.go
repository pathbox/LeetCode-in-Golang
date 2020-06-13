package Offer026

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return isContain(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B) // 这里看到区别了吧,当前节点与B进行开始前序遍历对比，要么A继续前序遍历，其当前节点的左右子树开始继续与B对比
}

func isContain(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}
	// 说明当前节点相等，继续前序对比左右子树
	return isContain(A.Left, B.Left) && isContain(A.Right, B.Right)
}
