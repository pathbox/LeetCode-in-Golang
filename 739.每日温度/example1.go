package LeetCode739

func dailyTemperatures(T []int) []int {
	if len(T) == 0 || len(T) == 1 {
		return T
	}
	//创建一个数据来保存结果
	result := make([]int, len(T), len(T))
	//创建栈
	stack := make([]int, 0)
	//遍历当前数组
	for k, v := range T {
		// 栈内有元素
		for len(stack) > 0 {
			// 获取栈顶元素
			//获取栈顶元素（T中数据的下标）数据的下标
			st := stack[len(stack)-1]
			if v > T[st] {
				//栈顶元素比较小，说明温度上升，则当前元素下标减去栈顶元素就是天数
				result[st] = k - st
				//栈顶元素弹出
				stack = stack[0 : len(stack)-1]
			} else {
				//栈顶元素比数据大，不做操作跳出循环，若不跳出会死循环
				break
			}
		}
		stack = append(stack, k)
	}
	return result
}
