package LeetCode1028

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/solution/1028-cong-xian-xu-bian-li-huan-yuan-er-cha-shu-zi-/
func recoverFromPreorder(S string) *TreeNode {
	ans := map[int]*TreeNode{-1: &TreeNode{}}
	addTree := func(v, p int) {
		ans[p] = &TreeNode{Val: v}
		if ans[p-1].Left == nil {
			ans[p-1].Left = ans[p]
		} else {
			ans[p-1].Right = ans[p]
		}
	}

	val, dep, hasVal := 0, 0, false
	for _, c := range S {
		if c != '-' {
			val = 10*val + int(c-'0')
			hasVal = true
		} else if hasVal {
			addTree(val, dep)
			val, dep, hasVal = 0, 1, false
		} else {
			dep++
		}
	}
	addTree(val, dep)
	return ans[0]
}
