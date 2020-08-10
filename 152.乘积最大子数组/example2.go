package LeetCode152

import "sort"

func maxProduct(nums []int) int {
	maxDP := make([]int, len(nums))
	minDP := make([]int, len(nums))
	maxDP[0] = nums[0]
	minDP[0] = nums[0]

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		maxDP[i], minDP[i] = numCompare(minDP[i-1]*nums[i], maxDP[i-1]*nums[i], nums[i])
		if max < maxDP[i] {
			max = maxDP[i]
		}
	}
	return max
}

func numCompare(a, b, c int) (max, min int) {
	s := []int{a, b, c}
	sort.Ints(s)
	return s[2], s[0]
}
