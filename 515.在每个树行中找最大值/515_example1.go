package LeetCode515

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

// 层次遍历
func largestValues(root *TreeNode) []int {
	res = make([]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		max := -1 << 63
		for i := 0; i < length; i++ {
			if max < queue[i].Val {
				max = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, max)
		queue = queue[length:] // 前length已经遍历过，不会进入到下一层遍历
	}
	return res
}
