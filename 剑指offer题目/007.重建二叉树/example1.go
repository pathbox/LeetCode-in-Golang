package Offer007

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/solution/er-cha-shu-de-qian-xu-bian-li-fen-zhi-si-xiang-by-/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}

	var index int
	for i := range inorder {
		if inorder[i] == preorder[0] {
			index = i // 找到根节点的索引
			break
		}
	}
	// 构造三个节点，三个节点为一个子树, 上面已经构造了root节点，下面是构造左右子树
	root.Left = buildTree(preorder[1:index+1], inorder[:index])   // 确定左子树
	root.Right = buildTree(preorder[index+1:], inorder[index+1:]) // 确定右子树
	return root
}
