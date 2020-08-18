package LeetCode581

import "math"

func findUnsortedSubarray(nums []int) int {
	res := 0
	L, R := -1, -1
	min, max := math.MaxInt64, math.MinInt64 //已遍历序列中的min,max
	//正向遍历，寻找右边界（最右的逆序）
	for i := 0; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			R = i
		}
	}
	//反向遍历，寻找左边界（最左的逆序）
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			L = i
		}
	}
	if R > L {
		res = R - L + 1
	}
	return res
}
