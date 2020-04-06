package LeetCode101

func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, root, root)
	for len(queue) >= 2 { // 不断循环读取队列
		t1 := queue[0]
		t2 := queue[1]
		queue = queue[2:]
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		// 将后续需要比较的节点放入队列
		queue = append(queue, t1.Left)
		queue = append(queue, t2.Right)
		queue = append(queue, t1.Right)
		queue = append(queue, t2.Left)
	}
	return true
}
