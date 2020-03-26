package LeetCode589

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res = make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for _, v := range root.Children {
		helper(v, res)
	}
}
