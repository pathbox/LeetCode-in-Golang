package LeetCode199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			length--
			if len(res) == level {
				res = append(res, queue[0].Val)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}

			queue = queue[1:]
		}
		level++
	}
	return res
}
