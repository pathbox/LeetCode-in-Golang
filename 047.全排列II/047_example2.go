package LeetCode047

import "sort"

// https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
func permuteUnique(nums []int) (result [][]int) {
	sort.Ints(nums)
	mp := make([]bool, len(nums))
	fill(nums, mp, []int{}, &result)
	return
}

func fill(nums []int, mp []bool, buf []int, result *[][]int) {
	n := len(nums)
	if n == len(buf) {
		cp := make([]int, len(buf))
		copy(cp, buf)
		*result = append(*result, cp)
		return
	}
	for i := 0; i < n; i++ {
		if mp[i] {
			continue
		}
		mp[i] = true
		v := nums[i]
		fill(nums, mp, append(buf, v), result)
		mp[i] = false
		// 去重
		for i+1 < n && nums[i+1] == v {
			i++
		}
	}
}
