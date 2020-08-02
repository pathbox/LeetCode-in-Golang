package LeetCode018

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := threeSum(nums[i+1:], target-nums[i])
		if tmp != nil {
			for _, t := range tmp {
				t = append(t, nums[i])
				res = append(res, t)
			}
		}
	}
	return res
}

func threeSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := twoSum(nums[i+1:], target-nums[i])
		if tmp != nil {
			for _, t := range tmp {
				t = append(t, nums[i])
				res = append(res, t)
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	has := make(map[int]bool)
	drop := make(map[int]bool)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if has[target-nums[i]] {
			if drop[target-nums[i]] == false {
				res = append(res, []int{target - nums[i], nums[i]})
				drop[target-nums[i]] = true
			}
		} else {
			has[nums[i]] = true
		}
	}
	return res
}

// O(N^3)
