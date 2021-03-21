package LeetCode032

// 具体做法是我们始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标。这样，新的有可能满足的结果是从最后一个没有被匹配的右括号下标之后的字符串开始的
func longestValidParentheses(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' { // 左括号 对应索引入栈
			stack = append(stack, i)
		} else { // 右括号
			stack = stack[:len(stack)-1] // 栈顶出栈
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1]) // 当前i-栈顶的索引
			}
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
