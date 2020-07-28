package LeetCode611

import "sort"

func triangleNumber(nums []int) int {
	// check if there is more than 3 nums
	if len(nums) < 3 {
		return 0
	}

	// first 排序操作
	sort.Ints(nums)

	ans := 0
	n := len(nums)

	for c := n - 1; c > 1; c-- {
		b := c - 1
		a := 0
		for a < b {
			if nums[a]+nums[b] > nums[c] {
				ans += b - a
				b--
			} else {
				a++
			}
		}
	}
	return ans
}
