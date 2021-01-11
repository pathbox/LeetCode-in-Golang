package LeetCode863

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
res分成两部分：
1、从target开始dfs距离k的节点
2、从vecRoot2Target[len-2]开始dfs距离为k-1的节点
*/
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	parent := make(map[*TreeNode]*TreeNode)
	recordFather(root, nil, parent)
	result := &[]int{}
	dfs(target, K, make(map[*TreeNode]struct{}), result, parent)
	return *result
}

// 因为不仅会从子树去遍历，还会从father节点去dfs，就有可能遇到重复遍历的情况，所以需要dup来过滤重复遍历的情况
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
		// 三种dfs情况，子树和从父节点
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
