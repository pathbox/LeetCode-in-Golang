package LeetCode257

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：每个节点只会被访问一次，因此时间复杂度为 O(N)O(N)，其中 NN 表示节点数目。
// 空间复杂度：O(N)O(N)。这里不考虑存储答案 paths 使用的空间，仅考虑额外的空间复杂度。额外的空间复杂度为递归时使用的栈空间，在最坏情况下，当二叉树中每个节点只有一个孩子节点时，递归的层数为 NN，此时空间复杂度为 O(N)O(N)。在最好情况下，当二叉树为平衡二叉树时，它的高度为 \log(N)log(N)，此时空间复杂度为 O(\log(N))O(log(N))。

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	s := ""
	if root == nil {
		return res
	}

	Run(&res, root, s)
	return res
}

func Run(res *[]string, root *TreeNode, s string) string {
	if root.Left == nil && root.Right == nil { // 叶子节点
		s += strconv.Itoa(root.Val)
		*res = append(*res, s)
		return ""
	}

	s += strconv.Itoa(root.Val) // 拼接数字
	s += "->"
	if root.Left != nil {
		Run(res, root.Left, s) // 继续左子树递归
	}
	if root.Right != nil {
		Run(res, root.Right, s) // 继续右子树递归
	}
	return s
}
