package max_distance

var res int

// 二叉树节点间的最大距离问题
func maxDistance(root *Node) int {
	return postOrder(root)
}

func posOrder(node *Node) int {
	if node == nil {
		res = 0
		return 0
	}
	lm := posOrder(node.Left)
	maxLeft = res
	rm := posOrder(node.Right)
	maxRight = res
	curMax := maxLeft + maxRight + 1
	res = max(maxLeft, maxRight) + 1
	return max(max(lm, rm), curMax)
}
