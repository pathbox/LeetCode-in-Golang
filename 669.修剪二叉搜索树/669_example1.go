package LeetCode669

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O(N)O(N)，其中 NN 是给定的树的全部节点。我们最多访问每个节点一次。

// 空间复杂度：O(N)O(N)，即使我们没有明确使用任何额外的内存，在最糟糕的情况下，我们递归调用的栈可能与节点数一样大

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > R { // 左子树更小，所以应该往左子树遍历,返回得到的node
		return trimBST(root.Left, L, R)
	}
	if root.Val < L { // 右子树更大，所以应该往右子树遍历,返回得到的node
		return trimBST(root.Right, L, R)
	}

	root.Left = trimBST(root.Left, L, R)   // 左子树 不断遍历Left
	root.Right = trimBST(root.Right, L, R) // 右子树  不断遍历Right
	return root
}
