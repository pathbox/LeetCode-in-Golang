package Offer031

/*
1，用go语言slice建一个辅助栈，来顺序添加压入序列
2，每次添加一个元素，去跟弹出序列比较，若不同，继续添加。若相同，则弹出，继续拿新建栈的栈顶元素去跟弹出序列去比较，循环即可。
*/

func validateStackSequences(pushed []int, popped []int) bool {
	// 建一个辅助栈
	stack := make([]int, 0)
	// 将push顺序压入辅助栈中，如果栈顶元素== pop序列中下一个出现的值，则弹出
	i := 0
	for _, value := range pushed {
		stack = append(stack, value)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] { // 栈顶元素和popped当前的值相等，则说明栈顶元素此时要出栈
			stack = stack[:len(stack)-1]
			i++
		}
	}

	// 总结判断
	if len(stack) == 0 { // 0 说明对应出栈完全成功
		return true
	} else {
		return false
	}
}
