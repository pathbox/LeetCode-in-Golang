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
	// 中顺序列找根节点，前序数组用于定位每个子树的根节点，中序数组根据这个根节点分成左右子树部分
	var root int
	for k, v := range inorder {
		if v == preorder[0] { // preorder[0] 是root，找到中序数组中root的index位置，以该位置分成左右子树两部分进行递归操作
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
		Left:  buildTree(preorder[1:root+1], inorder[0:root]), // inorder root索引值就是preorder[0]了，不用再来带到左子树中取一次使用
		Right: buildTree(preorder[root+1:], inorder[root+1:]),
	}
}
