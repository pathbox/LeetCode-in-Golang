package LeetCode027

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)

	return root
}

/*
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left, root.Right = root.Right, root.Left
    mirrorTree(root.Left)
    mirrorTree(root.Right)
    return root
}
*/
