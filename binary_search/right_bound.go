package binarysearch

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 锁定右边界，缩小右边界
			left = mid + 1
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
