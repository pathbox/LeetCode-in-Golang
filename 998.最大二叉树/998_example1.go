package LeetCode998

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if val < root.Val {
		insertRight(root, val)
		return root
	}
	return reRoot(root, val)
}

func insertRight(root *TreeNode, val int) {
	cursor := root
	for {
		if cursor.Right != nil {
			if cursor.Right.Val < val {
				cursor.Right = reRoot(cursor.Right, val)
				return
			}
			cursor = cursor.Right // 右子树遍历
		} else {
			break
		}
	}
	cursor.Right = newNode(val) // 加到右子树最后
	return
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func reRoot(root *TreeNode, val int) *TreeNode {
	newRoot := newNode(val)
	newRoot.Left = root
	return newRoot
}
