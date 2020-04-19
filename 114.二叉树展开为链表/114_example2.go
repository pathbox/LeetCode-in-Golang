package LeetCode114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	r := root.Right                        // 缓存当前root的右节点
	root.Right, root.Left = root.Left, nil // 当前节点的右子树被左子树替代，左子树置为空
	for root.Right != nil {                // 当前节点顺着右子树到最后一个右子树节点，
		root = root.Right
	}
	root.Right = r // 然后把之前缓存的root的右子树节点，连接到当前节点的右子树上
}
