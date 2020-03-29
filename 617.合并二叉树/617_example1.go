package LeetCode617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// 谁为nil，则返回另一方子树
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val // 前序遍历，先操作根节点，再遍历左子树和右子树
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// 时间复杂度：O(N)O(N)，其中 NN 是两棵树中节点个数的较小值。

// 空间复杂度：O(N)O(N)，在最坏情况下，会递归 NN 层，需要 O(N)O(N) 的栈空间。

// 树的节点并非只是一个单个节点，而是表示一个子树的节点，其下有可能还连这一个子树
