package LeetCode448

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)
	for _, value := range nums {
		if value < 0 {
			value = -value // 把小于0的值转为大于0
		}
		if nums[value-1] > 0 { // value-1为index进行判断
			nums[value-1] = -nums[value-1]
		}
	}

	for index, value := range nums {
		if value > 0 {
			result = append(result, index+1)
		}
	}
	return result
}
