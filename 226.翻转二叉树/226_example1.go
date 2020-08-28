package LeetCode226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1. 递归得到right节点
2. 递归得到left节点
3. 到root层，root的left赋值为上述递归得到的right，right为上述递归得到的left，返回当前root节点
树的递归是自底向上的
*/
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
