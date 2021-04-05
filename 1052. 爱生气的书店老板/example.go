package LeetCode1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	// 计算X天能够得到的最大额外收益
	total := 0 // total 是不使用技能时候的正常收益
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}

	increase := 0
	for i := 0; i < X; i++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrease := increase
	// 以X长度滑动窗口移动
	for i := X; i < n; i++ {
		increase = increase - customers[i-X]*grumpy[i-X] + customers[i]*grumpy[i] // i - X 其实是左指针位置， i是右指针
		maxIncrease = max(maxIncrease, increase)
	}
	return total + maxIncrease
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
