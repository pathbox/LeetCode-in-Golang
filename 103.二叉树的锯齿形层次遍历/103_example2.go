package LeetCode103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	left := true
	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		item := make([]int, len(curLevel))
		nextLevel := make([]*TreeNode, 0)
		for i, n := range curLevel {
			if left {
				item[i] = n.Val
			} else {
				item[len(curLevel)-i-1] = n.Val
			}
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}
		if left {
			left = false
		} else {
			left = true
		}
		result = append(result, item)
		curLevel = nextLevel
	}
	return result
}
