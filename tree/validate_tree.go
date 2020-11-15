package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// IsBinarySearchTree 校验是否是二叉搜索树
func IsBinarySearchTree(root *Node) bool {
	min := -1 << 63
	max := 1<<63 - 1
	res := dfs(root, min, max)
	return res
}

func dfs(root *Node, min, max int) bool {
	return root == nil || (min < root.Val && root.Val < max) && dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}

// IsCBT 是否是完全二叉树
func IsCBT(root *Node) bool {
	if root == nil {
		return true
	}
	queue := []*Node{root}
	leaf := false
	for len(queue) > 0 {
		// count := len(queue)
		// for i := 0; i < count; i++ {
		node := queue[0]
		l := node.Left
		r := node.Right
		if leaf && (l != nil || r != nil) || (l == nil && r != nil) { // 有右孩子，没左孩子 当前孩子不是左右都有，之后的节点必须都为叶子节点
			return false
		}
		if l != nil {
			queue = append(queue, l)
		} else if r != nil {
			queue = append(queue, r)
		} else {
			leaf = true // 表示此时的node是叶子节点
		}
		// }
	}
	return true
}
