package LeetCode590

type Node struct {
	Val      int
	Children []*Node
}

// 先加Children节点再加入根节点
func postorder(root *Node) []int {
	var res = make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for _, v := range root.Children { // 最后的子节点 Children是空的，for不会执行
		helper(v, res)
	}
	*res = append(*res, root.Val)
}
