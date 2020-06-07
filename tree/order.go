package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [3,9,20,null,null,15,7]
	//  [1,2,2,3,3,null,null,4,4]
	treeT := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	treeF := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	// 测试打印下前序遍历这两个二叉树
	fmt.Println("\nPreOrder treeT : ")
	PreOrder(treeT)
	fmt.Println("\nPreOrder treeF : ")
	PreOrder(treeF)

	// 测试打印下中序遍历这两个二叉树
	fmt.Println("\nInfixOrder treeT : ")
	InfixOrder(treeT)
	fmt.Println("\nInfixOrder treeF : ")
	InfixOrder(treeF)

	// 测试打印下后序遍历这两个二叉树
	fmt.Println("\nPostOrder treeT : ")
	PostOrder(treeT)
	fmt.Println("\nPostOrder treeF : ")
	PostOrder(treeF)

}

// PreOrder 先序遍历
func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%v; ", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

// InfixOrder 中序遍历
func InfixOrder(root *TreeNode) {
	if root == nil {
		return
	}

	InfixOrder(root.Left)
	fmt.Printf("%v; ", root.Val)
	InfixOrder(root.Right)
}

// PostOrder 后序遍历
func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}

	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%v; ", root.Val)
}
