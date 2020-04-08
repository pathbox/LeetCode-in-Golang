package LeetCode111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func minDepth(root *TreeNode) int {
	var level int
	if root == nil {
		return level
	}

	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		level++
		length := len(queue)
		for 0 < length {
			length--
			if queue[0].Left == nil && queue[0].Right == nil {
				return level // 找到第一个叶子节点，就是min depth
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left) // 把第一个的左子树继续入队列
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}
	return level
}
