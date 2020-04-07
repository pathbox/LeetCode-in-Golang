package LeetCode107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int
	var queue []TreeNode

	if root == nil {
		return ret
	}

	queue = append(queue, *root)
	for len(queue) > 0 {
		var tmp []int
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			tmp = append(tmp, queue[i].Val)

			if queue[i].Left != nil {
				queue = append(queue, *queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, *queue[i].Right)
			}
		}
		ret = append(ret, tmp)
		queue = queue[queueLen:]
	}

	lp := 0
	rp := len(ret) - 1
	for lp < rp {
		ret[lp], ret[rp] = ret[rp], ret[lp]
		lp++
		rp--
	}
	return ret
}
