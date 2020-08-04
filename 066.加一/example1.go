package LeetCode066

func plusOne(digits []int) []int {
	l := len(digits)

	if l == 0 {
		return []int{1}
	}

	for i := l - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits // 如果不是9 加完就返回
		} else {
			digits[i] = 0 // 是9 加完为0 在下一个循环中会进位
		}
	}
	// 如果还没返回，遍历了整个数组，则说明进位到了开头，需要补最高位1
	return append([]int{1}, digits...)
}
