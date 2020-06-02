package offer003

import "sort"

// 将数组排序后，再用双指针，前后移动进行比较, lastNum的方式可以只用一个指针

// 原地置换 O(n) 空间O(1)
func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	lastNum := nums[0]
	for _, num := range nums[1:] {
		if lastNum == num {
			break
		} else {
			lastNum = num // 把值赋给lastNum让它和下一位比较
		}
	}
	return lastNum
}
