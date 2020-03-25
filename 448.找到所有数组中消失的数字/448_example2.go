package LeetCode448

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)
	for _, value := range nums {
		if value < 0 {
			value = -value // 把小于0的值转为大于0
		}
		if nums[value-1] > 0 { // value-1为index进行判断,索引值从0开始
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

// 一个萝卜一个坑，如果这个坑没有被标记过，说明这个坑没被用过，也就是本来在这个坑上的数没有在数组中出现过，也就是消逝的数
