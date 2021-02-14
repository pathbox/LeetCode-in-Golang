package LeetCode090

import "sort"

var (
	res [][]int
)

func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	backtrack([]int{}, nums, 0)
	return res
}

func backtrack(temp, nums []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp) // copy一份
	res = append(res, tmp)

	for i := start; i < len(nums); i++ { // start很关键，这样不会导致[2,1]这种重复的情况，2只会和比它大的数组合。 i == start的是可以重复的
		if i != start && nums[i] == nums[i-1] {
			continue
		}

		temp = append(temp, nums[i]) // i
		backtrack(temp, nums, i+1)   // i+1 temp进入的数只会从后面选，不会从倒回从前选，避免了重复情况

		// 通过删掉最后一个元素实现回溯， 才会有 回溯 1， 再回溯 2，最后回溯 3
		temp = temp[:len(temp)-1]
	}
}
