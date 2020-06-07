package Offer011

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	min := -1
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			min = numbers[i]
		}
	}

	if min != -1 { // 说明数组没有旋转
		min = numbers[0]
	}

	return min
}
