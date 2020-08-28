package Offer007

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/solution/er-cha-shu-de-qian-xu-bian-li-fen-zhi-si-xiang-by-/
/*
通过preorder[0]得到某个root节点
定位inorder root节点的位置
现在知道了root节点在preorder和inorder的位置，再继续递归左子树和右子树，下一层的递归通过preorder和inorder排除已经找到的root节点
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}

	var index int
	for i := range inorder {
		if inorder[i] == preorder[0] { // 每一层的preorder[0]，inorder[index]是根节点，以此继续递归左子树的preorder和inorder参数传入和右子树的
			index = i // 找到根节点的索引
			break
		}
	}
	// 构造三个节点，三个节点为一个子树, 上面已经构造了root节点，下面是构造左右子树
	// preorder[0]是根节点,inorder[index]是根节点 已经在上面使用
	root.Left = buildTree(preorder[1:index+1], inorder[:index])   // 确定左子树
	root.Right = buildTree(preorder[index+1:], inorder[index+1:]) // 确定右子树
	return root
}
