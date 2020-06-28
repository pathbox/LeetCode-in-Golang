package Offer055

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//层序遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var (
		queue []*TreeNode
		res   int
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		lens := len(queue)
		for i := 0; i < lens; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[lens:]
		res++
	}
	return res
}
