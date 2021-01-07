package LeetCode1028

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/solution/shou-hui-tu-jie-fei-di-gui-fa-zhong-gou-chu-er-cha/
func recoverFromPreorder(S string) *TreeNode {
	path, pos := []*TreeNode{}, 0
	for pos < len(S) {
		level := 0 // level初始值
		// pos是不断的往前增加的
		for S[pos] == '-' {
			level++
			pos++
		}
		value := 0
		// 计算当前value值
		for ; pos < len(S) && S[pos] >= '0' && S[pos] <= '9'; pos++ {
			value = value*10 + int(S[pos]-'0')
		}
		node := &TreeNode{Val: value}
		if level == len(path) { // 此时栈的top就是当前遍历节点的父节点
			if len(path) > 0 {
				path[len(path)-1].Left = node
			}
		} else {
			path = path[:level] // 大于level的节点全部删除“出栈”,找到path[:level]的位置,node是在后面遍历到的节点，只能是右节点
			path[len(path)-1].Right = node
		}
		path = append(path, node)
	}
	return path[0]
}
