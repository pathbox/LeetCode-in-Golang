package LeetCode088

import "sort"

// 先合并到nums1 然后再排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}
