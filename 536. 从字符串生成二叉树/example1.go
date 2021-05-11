package LeetCode536

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func str2tree(s string) *TreeNode {
	if 0 == len(s) {
		return nil
	}

	// 对于A(B)(C)，即有root.val(root.left.val)(root.right.val)
	// 以上符合前序遍历树，使用栈存储节点

	stack := make([]*TreeNode, 0)

	for i := 0; i < len(s); i++ {
		// 遇到)，弹出顶层元素，保持栈顶为根节点，非叶子节点,最后需要返回根节点就可以
		if s[i] == ')' {
			stack = stack[:len(stack)-1]
		} else if (s[i] >= '0' && s[i] <= '9') || s[i] == '-' {
			// 遇到整数
			start := i
			// 负数：如果i是负号且i+1是数值，接着往后知道找到(或)或结束的位置
			for i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
				i++
			}
			// 取数值（可能是整数或负数）
			val, _ := strconv.Atoi(s[start : i+1])
			node := &TreeNode{Val: val}

			// 如果栈不为空，当前节点必是栈顶元素（根节点）的左或右叶子节点
			if len(stack) != 0 {
				topNode := stack[len(stack)-1]
				if topNode.Left == nil {
					topNode.Left = node
				} else {
					topNode.Right = node
				}
			}
			// 把当前节点入栈
			stack = append(stack, node)
		}
	}

	// 经过多次判断)符号的出栈，栈顶必是根节点
	return stack[len(stack)-1]
}
