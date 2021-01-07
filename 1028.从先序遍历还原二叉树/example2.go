package LeetCode1028

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/solution/1028-cong-xian-xu-bian-li-huan-yuan-er-cha-shu-zi-/
/*
由先序遍历的特性可以知道，上一个遍历到的比自己深度浅1的节点必为自己的父节点，所以直接把当前节点加到这父节点下便可以把树连接起来了
*/
func recoverFromPreorder(S string) *TreeNode {
	ans := map[int]*TreeNode{-1: &TreeNode{}}
	addTree := func(v, p int) {
		ans[p] = &TreeNode{Val: v}
		if ans[p-1].Left == nil { // 先选择加入左节点
			ans[p-1].Left = ans[p]
		} else {
			ans[p-1].Right = ans[p]
		}
	}

	val, dep, hasVal := 0, 0, false
	for _, c := range S { // 遍历S字符串
		if c != '-' { // 先判断是否是数值
			val = 10*val + int(c-'0')
			hasVal = true
		} else if hasVal { // 如果有合法的数值，加入到树中
			addTree(val, dep)
			val, dep, hasVal = 0, 1, false // 每加入一次树中，树的dep变为1,下一个节点至少有父节点了，所以dep是从1开始初始值
		} else { // c 是 - 表示深度需要++
			dep++
		}
	}
	addTree(val, dep)
	return ans[0]
}
