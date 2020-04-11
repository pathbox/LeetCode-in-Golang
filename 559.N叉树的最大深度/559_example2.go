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
		length := len(queue) // 当前层的节点数
		for length > 0 {     // 将当前这层的所有node的Children入queue
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

// 每次进入一层，遍历这一层的所有节点数，将其Children节点入队列，将当前这层parent节点处队列，这一层level++
// 然后继续对Children节点进行相同的操作，直到所有节点没有Children为止，说明是到了最长的叶子节点位置
