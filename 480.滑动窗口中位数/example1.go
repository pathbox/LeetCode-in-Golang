package LeetCode480

import "sort"

func medianSlidingWindow(nums []int, k int) []float64 {
	knums := make([]int, k)
	copy(knums, nums[:k])
	sort.Ints(knums)
	ret := make([]float64, len(nums)-k+1)
	ret[0] = calMid(knums, k)
	for i := k; i < len(nums); i++ {
		// 用二分查找+插入排序 使得knums始终是排序数组窗口
		delIdx := binarySearch(knums, nums[i-k])
		knums = append(knums[:delIdx], knums[delIdx+1:]...)
		addIdx := binarySearch(knums, nums[i])
		last := append([]int{nums[i]}, knums[addIdx:]...)
		knums = append(knums[:addIdx], last...)
		ret[i-k+1] = calMid(knums, k)
	}
	return ret
}

func binarySearch(nums []int, val int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>2
		if nums[mid] == val {
			return mid
		} else if nums[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func calMid(nums []int, k int) float64 {
	n1, n2 := 0, 0
	if k%2 == 0 {
		n1, n2 = nums[k/2-1], nums[k/2]
	} else {
		n1, n2 = nums[k/2], nums[k/2]
	}
	return float64(n1+n2) / float64(2)
}
