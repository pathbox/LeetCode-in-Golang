package LeetCode098

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue) // 获得这一层的所有节点数量
		res = append(res, []int{})
		for counter > 0 { // 开始遍历这一层，counter = 0 表示一层遍历完
			counter--
			// 每个节点的左子树 和右子树入队列，在下一层进行遍历
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val) // 将当前层的节点输入的这一层的数组中
			queue = queue[1:]
		}
		level++
	}
	return res
}
