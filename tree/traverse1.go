package tree

import "fmt"

type TreeNode struct {
	ID    int
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Println(root.Val)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

// 中序遍历
func InOrder(root *TreeNode) {
	if root != nil {
		InOrder(root.Left)
		// root.Left = nOrder(root.Left)
		fmt.Println(root.Left)
		InOrder(root.Right)
	}
}

// 后序遍历
func PostOrder(root *TreeNode) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		fmt.Printf("%d ", root.Val)
	}
}
