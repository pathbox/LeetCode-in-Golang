package LeetCode103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 中顺序列找根节点
	var root int
	for k, v := range inorder {
		if v == preorder[0] {
			root = k
			break
		}
	}

	// 左右子树归类
	// pre_left, pre_right := preorder[1: root+1], preorder[root+1:]
	// in_left, in_right := inorder[0: root], inorder[root+1:]
	// 左右子树递归
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:root+1], inorder[0:root]),
		Right: buildTree(preorder[root+1:], inorder[root+1:]),
	}
}
