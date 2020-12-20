### 二分查找的三种情况

```go
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 锁定左边界，往左缩小区域 缩小搜索区间的上界
			right = mid - 1
		}
	}

	// 最后检查left越界情况,或者没有匹配
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 锁定右边界，往右缩小区域 缩小搜索区间的下界
			left = mid + 1
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

```