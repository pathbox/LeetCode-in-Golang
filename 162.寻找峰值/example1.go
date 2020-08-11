package LeetCode162

// https://leetcode-cn.com/problems/find-peak-element/solution/hua-jie-suan-fa-162-xun-zhao-feng-zhi-by-guanpengc/
func findPeakElement(nums []int) int {
	var (
		left, mid int
		right     = len(nums) - 1
	)

	for left < right {
		mid = (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
