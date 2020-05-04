package LeetCode1325

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 它的子节点一定先于它被操作。这其实也就是二叉树的后序遍历
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target { // 后序遍历到子节点,Val和target匹配，返回nil，则表示当前子节点被删除，不会返回给上层的父节点
		return nil
	}
	return root
}
