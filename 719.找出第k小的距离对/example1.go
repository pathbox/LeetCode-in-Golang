package LeetCode719

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	low, high := 0, nums[len(nums)-1]-nums[0]
	for low < high {
		mid := low + (high-low)>>1 // mid 是dis的二分中间
		if check(nums, mid, k) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func check(nums []int, dis int, k int) bool {
	low, count := 0, 0 // count 是 <= dis 的区间大小
	for high, _ := range nums {
		if nums[high]-nums[low] > dis {
			low++
		}
		count += high - low
	}
	return count >= k
}
