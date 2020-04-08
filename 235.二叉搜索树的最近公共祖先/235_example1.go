package LeetCode235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 可以发现，如果满足 p.Val <= 当前节点.Val <= q.Val，那么说明该节点就是 p 和 q 的最近公共祖先，返回当前节点的地址。如果不符合，就根据 p 和 q 的值判断去左子树找还是去右子树找。

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	if root == nil || p.Val <= root.Val && root.Val <= q.Val {
		return root // 找到该节点
	} else if q.Val <= root.Val { // root值太大 从左子树中寻找
		return lowestCommonAncestor(root.Left, p, q)
	} else { // 从右子树中寻找
		return lowestCommonAncestor(root.Right, p, q)
	}
}
