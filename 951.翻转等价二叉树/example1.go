package LeetCode951

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1、如果root1，root2节点的左右子树节点值相等，那么，判断root1，root2左子树是否可以翻转，判断root1，root2右子树是否可以翻转
// 2、如果root1，root2节点的左右子树节点值可以通过交换的方式使得相等，那么交换，然后继续步骤一
// 3、如果1，2的两个如果均不满足，那么返回false

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if (root1 == nil && root2 != nil) || (root2 == nil && root1 != nil) {
		return false
	}
	if root1.Val == root2.Val { // 包括进行一次翻转和0次翻转的情况，所以对应节点不翻转，而是相等，表示进行0次翻转，也满足条件
		return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
			(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
	}

	return false
}
