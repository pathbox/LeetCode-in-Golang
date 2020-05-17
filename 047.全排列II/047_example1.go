package LeetCode047

import "sort"

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	backtrack(nums, nil, &result)
	return result
}

func backtrack(nums, track []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, track)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		numsCopy := append([]int{}, nums...)
		trackCopy := append([]int{}, track...)
		backtrack(append(numsCopy[:i], numsCopy[i+1:]...), append(trackCopy, nums[i]), result)
	}
}
