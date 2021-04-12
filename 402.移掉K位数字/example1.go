package LeetCode402

/*
让我们从一个简单的例子开始。给定一个数字序列，例如 425，如果要求我们只删除一个数字，那么从左到右，我们有 4、2 和 5 三个选择。我们将每一个数字和它的左邻居进行比较。从 2 开始，小于它的左邻居 4。则我们应该去掉数字 4。如果不这么做，则随后无论做什么，都不会得到最小数。

如果我们保留数字 4，那么所有可能的组合都是以数字 4（即 42，45）开头的。相反，如果去掉 4，留下 2，我们得到的是以 2 开头的组合（即 25），这明显小于任何留下数字 4 的组合。
利用栈，将栈中比其大的数剔除，然后自己入栈
*/
func removeKdigits(num string, k int) string {
	var stack []uint8               // 使用 slice 作为栈
	var result string               // 存储最终结果字符串
	for i := 0; i < len(num); i++ { // 从最高位扫描数字字符串num
		number := num[i] - '0'
		for len(stack) != 0 && stack[len(stack)-1] > number && k > 0 {
			stack = stack[:len(stack)-1]
			k-- // 剔除一个数字，k 减一
		}
		// 条件：如果数字不为空，或者是栈不为空 将当前number入栈
		if number != 0 || len(stack) != 0 {
			stack = append(stack, number)
		}
	}
	// 条件：如果栈不空且还能删除数字
	for len(stack) != 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	// 将栈中元素转化为字符串
	for _, v := range stack {
		result += string('0' + v)
	}
	// 注意为空情况
	if result == "" {
		return "0"
	}
	return result
}
