package LeetCode154

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>2
		// 只和right边比较即可
		if nums[mid] < nums[right] {
			right = mid // 最小值可能是mid，所以区域可以从mid开始
		} else if nums[mid] > nums[right] {
			left = mid + 1 // left=mid 可以吗？不可以，因为已经知道最小值不可能是mid，right还更小
		} else if nums[mid] == nums[right] {
			right--
		}
	}
	return nums[left]
}
