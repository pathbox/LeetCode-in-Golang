package LeetCode101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left) // 三种情况，遍历到的节点值相同， 左子树节点等于右子树节点， 右子树节点等于左子树节点
}
