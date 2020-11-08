package Offer036

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var queue []*TreeNode

func treeToDoublyList(root *TreeNode) *TreeNode {
	inOrder(root)
	if len(queue) == 0 {
		return root
	}
	pre := root
	cur := &TreeNode{}
	for len(queue) > 0 {
		cur = queue[0]
		pre.Right = cur
		cur.Left = pre
		pre = cur
		queue = queue[1:]
	}
	pre.Right = nil
	return root
}

func inOrder(root *TreeNode) {
	inOrder(root.Left)
	queue = append(queue, root)
	inOrder(root.Right)
}
