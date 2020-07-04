package LeetCode033

func search(nums []int, target int) int {
	lo, lh := 0, len(nums)-1

	for lo <= lh {
		mid := lo + (lh-lo)/2
		if nums[mid] == target {
			return mid
		}

		// 先根据 nums[mid] 与 nums[lo] 的关系判断 mid 是在左段还是右段
		if nums[mid] >= nums[lo] {
			if target >= nums[lo] && target < nums[mid] {
				lh = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[lh] {
				lo = mid + 1
			} else {
				lh = mid - 1
			}
		}
	}
	return -1 // 没有找到匹配
}
