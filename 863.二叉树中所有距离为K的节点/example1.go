package LeetCode863

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	parent := make(map[*TreeNode]*TreeNode)
	recordFather(root, nil, parent)
	result := &[]int{}
	dfs(target, K, make(map[*TreeNode]struct{}), result, parent)
	return *result
}

func dfs(node *TreeNode, K int, dup map[*TreeNode]struct{}, result *[]int, parent map[*TreeNode]*TreeNode) {
	if node == nil {
		return
	}
	if _, exist := dup[node]; exist {
		return
	}
	dup[node] = struct{}{}
	if K == 0 {
		*result = append(*result, node.Val)
	} else {
		dfs(node.Left, K-1, dup, result, parent)
		dfs(node.Right, K-1, dup, result, parent)
		dfs(parent[node], K-1, dup, result, parent)
	}
}

func recordFather(root *TreeNode, father *TreeNode, parent map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}
	parent[root] = father
	recordFather(root.Left, root, parent)
	recordFather(root.Right, root, parent)
}
