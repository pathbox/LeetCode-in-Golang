package LeetCode081

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/solution/yi-wen-jie-jue-4-dao-sou-suo-xuan-zhuan-pai-xu-s-3/
// 这道题是 33 题的升级版，元素可以重复。当 nums[left] == nums[mid] 时，无法判断 target 位于左侧还是右侧，此时无法缩小区间，退化为顺序查找

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			return true
		}
		// 处理重复元素
		if nums[mid] == nums[left] {
			left++
		} else if nums[mid] == nums[right] {
			right--
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] > nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
