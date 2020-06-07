package Offer011

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		m := (left + right) / 2
		if numbers[left] < numbers[right] { // 如果这部分数组没有旋转过,则说明最小数的索引就是left
			return numbers[left]
		}

		if numbers[m] > numbers[right] {
			left = m + 1 // 说明最小值是在右边区域
		} else if numbers[m] < numbers[right] { // 再从左边区域寻找
			right = m
		} else {
			right-- // numbers[m] == numbers[right]  缩小右侧区域
		}
	}
	return numbers[left]
}
