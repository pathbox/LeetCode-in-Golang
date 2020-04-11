package LeetCode559

type Node struct {
	Val      int
	Children []*Node
}

// dfs
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, n := range root.Children {
		max = maxInt(max, maxDepth(n))
	}
	return max + 1
}

func maxInt(a, b int) int {
	if a < b {
		a = b
	}
	return a
}
