package main

import "fmt"

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func main() {
	var node TreeNode
	node = TreeNode{1, nil, nil}
	node.left = &TreeNode{2, nil, nil}
	node.right = &TreeNode{3, nil, nil}
	fmt.Println(inorder(&node))
}

func inorder(root *TreeNode) (res []int) {
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.value)
		root = root.right
	}
	return res
}

func preOrder(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.value)
			stack = append(stack, root)
			root = root.left
		}
		root = stack[len(stack)-1].right
		stack = stack[:len(stack)-1]
	}
	return res
}

func postOrder(root *TreeNode) (res []int) {
	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.right == nil || root.right == prev {
			res = append(res, root.value)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.right
		}
	}
	return res
}
