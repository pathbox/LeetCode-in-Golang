package LeetCode117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	switch {
	case root == nil:
		return nil
	case root.Left == nil && root.Right == nil:
		return root
	case root.Left == nil:
		root.Right.Next = next(root.Next)
		connect(root.Right)

	case root.Right == nil:
		root.Left.Next = next(root.Next)
		connect(root.Left)

	default:
		root.Left.Next = root.Right
		root.Right.Next = next(root.Next)
		// 需要先递归右子树
		connect(root.Right)
		connect(root.Left)
	}
	return root
}

func next(root *Node) *Node {
	if root == nil {
		return nil
	}

	for ; root != nil; root = root.Next {
		if root.Left != nil {
			return root.Left
		}

		if root.Right != nil {
			return root.Right
		}
	}
	return nil
}
