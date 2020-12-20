package LeetCode085

// https://leetcode-cn.com/problems/maximal-rectangle/solution/golang-zhuan-hua-wei-84ti-dan-diao-zhan-fa-by-bloo/

func maximalRectangle(matrix [][]byte) int {
	rowNum := len(matrix)
	if rowNum == 0 {
		return 0
	}
	colNum := len(matrix[0])
	heights := make([]int, colNum)
	var result int
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			if matrix[row][col] == '1' {
				heights[col] += 1
			} else {
				heights[col] = 0
			}
		}
		//调用上一题的解法，更新函数
		result = max(result, largestRectangleArea(heights))
	}
	return result
}

func largestRectangleArea(heights []int) int {
	// 加入哨兵值，便于原先heights中的最后位置的值弹出 左右两头哨兵
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
		tIndex := len(stack) - 1
		for len(stack) != 0 && heights[stack[tIndex]] > v {
			// pop栈顶元素
			topIndex := stack[tIndex]
			stack = stack[:tIndex]
			// 面积计算公式为：当前栈顶高度 * (左右两边的坐标减去1)
			result = max(result, heights[topIndex]*(k-stack[tIndex]-1))
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
