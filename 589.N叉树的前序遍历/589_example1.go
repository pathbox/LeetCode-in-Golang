package LeetCode589

type Node struct {
	Val      int
	Children []*Node
}

// 先加入root节点，再加入Children节点
func preorder(root *Node) []int {
	var res = make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)     // 先把root节点加入到数组
	for _, v := range root.Children { // 不断的遍历每个节点的 Children
		helper(v, res)
	}
}
