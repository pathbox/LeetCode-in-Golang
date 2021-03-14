package LeetCode032

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
