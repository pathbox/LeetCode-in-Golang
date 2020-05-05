package LeetCode1145

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	arrs := make([]int, 3)
	arrs[0], arrs[1], _ = find(root, x)
	arrs[2] = n - arrs[0] - arrs[1] - 1
	sort.Ints(arrs)
	return arrs[2] > arrs[0]+arrs[1]+1

}

func find(root *TreeNode, x int) (left, right int, isFind bool) {
	if root == nil {
		return 0, 0, false
	}
	if root.Val == x {
		return rootCount(root.Left), rootCount(root.Right), true

	}
	left, right, isFind = find(root.Left, x)
	if isFind {
		return
	}
	left, right, isFind = find(root.Right, x)
	return
}

// 计算root这个树的节点数
func rootCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + rootCount(root.Left) + rootCount(root.Right)
}
