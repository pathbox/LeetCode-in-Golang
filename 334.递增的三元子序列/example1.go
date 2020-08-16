package LeetCode334

import "math"

// 利用双指针
// 遍历查找下一个元素是否有更小元素，如果就更新，没有那么表示已达成三个递增序列
func increasingTriplet(nums []int) bool {
	m1, m2 := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if m1 >= v {
			m1 = v
		} else if m2 >= v {
			m2 = v
		} else {
			return true
		}
	}
	return false
}
