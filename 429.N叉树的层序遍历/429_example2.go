package LeetCode429

type Node struct {
	Val      int
	Children []*Node
}

var res [][]int

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	dfs(root, 0, &res)
	return res
}

//把每一层的level代入进来
func dfs(root *Node, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{}) // 加一个新的空层
	}
	(*res)[level] = append((*res)[level], root.Val) // 将当前节点的val加入到对应的level那一层
	for _, node := range root.Children {
		dfs(node, level+1, res)
	}

}
