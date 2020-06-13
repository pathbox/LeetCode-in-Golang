package Offer026

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	var ret bool

	if A.Val == B.Val {
		ret = helper(A, B)
	}
	//ret == false，说明 B 的根节点不在当前 A 树顶中，进入 A 的左子树进行递归查找
	if !ret {
		ret = isSubStructure(A.Left, B)
	}
	//当 B 的根节点不在当前 A 树顶和左子树中，进入 A 的右子树进行递归查找
	if !ret {
		ret = isSubStructure(A.Right, B)
	}
	return ret
	//利用 || 的短路特性可写成
	//return helper(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
}

//helper 校验 B 是否与 A 的一个子树拥有相同的结构和节点值
func helper(a, b *TreeNode) bool {
	if b == nil { // 说明B遍历完了，找到对应子树结构
		return true
	}
	if a == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}
	//a.Val == b.Val 递归校验 A B 左子树和右子树的结构和节点是否相同
	return helper(a.Left, b.Left) && helper(a.Right, b.Right)
}
