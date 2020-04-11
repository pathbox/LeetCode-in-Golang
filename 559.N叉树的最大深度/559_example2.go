package LeetCode559

type Node struct {
	Val      int
	Children []*Node
}

// bfs
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var queue = []*Node{root}
	var level int
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 { // 将当前这层的所有node的Children入queue
			length--
			for _, n := range queue[0].Children {
				queue = append(queue, n)
			}
			queue = queue[1:] // 当前这层的所有node都出队列
		}

		level++ // 这一层 加一
	}
	return level
}
