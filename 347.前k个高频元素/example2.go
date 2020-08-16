package LeetCode347

import "sort"

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	// 统计每个数字出现的次数
	rec := make(map[int]int, len(nums))
	for _, n := range nums {
		rec[n]++
	}
	// 对出现次数进行排序
	counts := make([]int, 0, len(rec))
	for _, c := range rec {
		counts = append(counts, c)
	}
	sort.Ints(counts)
	// min是前k个高频数字的底线
	min := counts[len(counts)-k]
	// 收集所有不低于底线的数字
	for n, c := range rec {
		if c >= min {
			res = append(res, n)
		}
	}
	return res
}

/* 桶排序法 时间复杂度：O(n)，n表示数组的长度。
首先，遍历一遍数组统计元素的频率，这一系列操作的时间复杂度是O(n)
桶的数量为n+1，所以桶排序的时间复杂度是O(n)
空间复杂度：O(n)
*/
