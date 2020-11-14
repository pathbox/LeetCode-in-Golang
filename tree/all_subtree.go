package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {

}

var roots []*Node
var routes []string

func allSubTree(root *Node) {
	if root == nil {
		return
	}
	helper(root)
}

// 递归是自底向上的
func helper(root *Node) {
	if root == nil {
		return nil
	}
	route := fmt.Sprintf("%s#%s#%s", root.Val, helper(root.Left), helper(root.Right)) // 先序遍历
	roots = append(roots, root)
	routes = append(routes, route) // 每个子树结构的前序打印字符串
	return route
}
