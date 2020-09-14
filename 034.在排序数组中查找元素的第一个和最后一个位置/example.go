package LeetCode034

func searchRange(nums []int, target int) []int {
	first, last := leftBound(nums, target), rightBound(nums, target)
	return []int{first, last}
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 搜索区间为 [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 处理下边界问题
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 搜索区间为 [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 处理下边界问题
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
