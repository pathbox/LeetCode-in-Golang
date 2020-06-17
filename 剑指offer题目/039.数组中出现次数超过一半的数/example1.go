package Offer039

import "sort"

// 排序 因为重复值超过了一半 肯定中间值也是
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
