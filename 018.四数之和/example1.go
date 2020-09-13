package LeetCode018

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r3 := threeSum(nums[i+1:], target-nums[i])
		for j := 0; j < len(r3); j++ {
			res = append(res, append(r3[j], nums[i])) // 三数之和的结果和当前nums[i]构成一个结果
		}
	}
	return res
}

func threeSum(nums []int, target int) [][]int {
	var buf [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		start, end := i+1, len(nums)-1
		for start < end {
			if start > i+1 && start < len(nums) && nums[start] == nums[start-1] { // start过滤重复
				start++
				continue
			}
			if end < len(nums)-1 && end >= start && nums[end] == nums[end+1] { // end过滤重复
				end--
				continue
			}
			if nums[i]+nums[start]+nums[end] > target {
				end--
			} else if nums[i]+nums[start]+nums[end] < target {
				start++
			} else {
				buf = append(buf, []int{nums[i], nums[start], nums[end]}) // 满足一种情况
				start++
				end--
			}
		}
	}
	return buf
}
