package LeetCode078

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	item := make([]int, 0)

	result = append(result, item)
	generate(0, nums, &item, &result)
	return result
}

func generate(i int, nums []int, item *[]int, result *[][]int) {
	if i >= len(nums) {
		return
	}
	*item = append(*item, nums[i])
	temp := make([]int, len(*item))
	for i, v := range *item {
		temp[i] = v
	}
	*result = append(*result, temp)
	generate(i+1, nums, item, result)
	*item = (*item)[:len(*item)-1] // 回溯
	generate(i+1, nums, item, result)
	return
}
