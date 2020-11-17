package LeetCode1714

import "sort"

// 排序
func smallestK(arr []int, k int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[:k]
}

// O(nlogn)
