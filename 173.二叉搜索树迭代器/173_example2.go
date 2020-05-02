package LeetCode173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	values []int
	index  int
}

func Constructor(root *TreeNode) BSTIterator {
	vals := make([]int, 0)
	inorder(root, &vals)
	return BSTIterator{
		values: vals,
		index:  0,
	}
}

func (bi *BSTIterator) Next() int {
	val := bi.values[bi.index] // 返回当前index的值
	bi.index++
	return val
}

func (bi *BSTIterator) HasNext() bool {
	return bi.index < len(b.values) // 遍历到底了
}

func inorder(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, values)
	*values = append(*values, root.Val) // 中序遍历，将val存入values数组，中序遍历二叉搜索树，values数组是递增的
	inorder(root.Right, values)
}
