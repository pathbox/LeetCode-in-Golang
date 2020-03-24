package LeetCode226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 不断递归到底部
	right := invertTree(root.Right) // 右子树不断的往下递归
	left := invertTree(root.Left)   // 左子树不断的往下递归

	// 交换
	root.Right = left
	root.Left = right

	// root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

// 右子树不断的往下递归,当递归到最后的右子树是nil，返回到上一层，不断递归左子树，直到 左子树为空，开始往上返回，交换此时root的左右子树， 相当于是自底向上交换
