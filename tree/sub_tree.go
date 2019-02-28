package sub_tree

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

func HasSubtree(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	// 判断B是否为A的子树(子结构)
	return isSubTree(root1, root2)
}

func isSubTree(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil { // A 为母树，节点比root2先递归完说明root1比root2小，root2不可能为root1的子树
		return false
	}
	// 以root1为root2的根节点，递归判断子结构是否成立
	if judge(root1, root2) {
		return true
	} else {
		return isSubTree(root1.left, root2) || isSubTree(root1.right, root2) // root1往左子树方向走 或 root2往右子树方向走
	}
}

func judge(root1 *TreeNode, root2 *TreeNode) bool {
	if root2 == nil { // root2 子树节点全判断完了，说明root2 是root1的子树
		return true
	}
	if root1 == nil {
		return false // A 为母树，节点比root2先递归完说明root1比root2小，root2不可能为root1的子树
	}
	if root1.val == root2.val {
		return judge(root1.left, right.left) && judge(root1.right, root2.right)
	}

	return false
}
