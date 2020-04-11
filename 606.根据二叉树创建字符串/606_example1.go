package LeetCode606

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil { // 情况1 t是叶子节点
		return strconv.Itoa(t.Val)
	}
	if t.Right == nil { // 情况2 右子树是nil  右子树不用输出空()
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	} // 情况3，4  遍历左右子树，左子树需要输出空()
	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")(" + tree2str(t.Right) + ")"
}
