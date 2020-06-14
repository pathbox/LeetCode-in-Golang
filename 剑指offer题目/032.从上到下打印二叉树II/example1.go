package Offer032

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS 广度优先遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
		queue = queue[length:]
	}
	return res
}
