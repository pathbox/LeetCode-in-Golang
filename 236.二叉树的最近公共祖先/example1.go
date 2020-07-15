package LeetCode236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil { // 找到了一个节点，返回那个节点，或者找到了叶子节点，直接返回
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q) // 如果该子树下包含要找的节点，那么返回该节点地址，否则返回 nil
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil { // // 如果返回的结果left和right都不为空，说明左右子树上各自找到p或q节点.因为树的递归返回是从底向上的，所以第一个满足这个条件返回的root，就是深度最深的公共祖先节点
		return root
	} else if l != nil {
		return l
	} else if r != nil {
		return r
	}
	return nil
}
