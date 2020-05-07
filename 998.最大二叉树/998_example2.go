package LeetCode998

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 主要理解题意，如果val > root.Val, 直接使之成为根,原树成为其左子树，反之，插入到右子树中
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if val > root.Val {
		return &TreeNode{val, root, nil}
	}
	root.Right = insertIntoMaxTree(root.Right, val) // 继续在右子树递归,会对右子树进行操作
	return root
}
