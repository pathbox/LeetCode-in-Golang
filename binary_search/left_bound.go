package binarysearch

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 锁定左边界，缩小右边界
			right = mid - 1
		}
	}

	// 最后检查left越界情况,或者没有匹配
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}
