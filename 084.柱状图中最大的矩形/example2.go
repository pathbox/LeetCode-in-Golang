package LeetCode084

func largestRectangleArea(heights []int) int {
	// 加入哨兵值，便于原先heights中的最后位置的值弹出
	heights = append(heights, 0)
	// 压入哨兵值，便于heights打头的数组进行操作
	heights = append([]int{0}, heights...)
	var stack []int
	var result int
	for k, v := range heights {
		// 栈里面后面比前面大的时候才压入，相当于顺序压入
		// 当前值比栈顶的值小的时候，相当于两个比栈顶小的值把栈顶位置的数卡在中间
		// 比如5，6，2，栈顶数为6
		// 此时可以计算栈顶6围成的矩形面积
		for len(stack) != 0 && heights[stack[len(stack)-1]] > v {
			// pop 栈顶元素
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 面积计算公式为：当前栈顶高度 * (左右两边的坐标减去1)
			result = max(result, heights[topIndex]*(k-stack[len(stack)-1]-1))
		}
		// 栈前面都为比当前值小的时候，无法将栈顶值卡在中间了
		// 此时压入当前坐标
		stack = append(stack, k)
	}
	return result
}

func max(a int, nums ...int) int {
	result := a
	for _, v := range nums {
		if v > result {
			result = v
		}
	}
	return result
}
