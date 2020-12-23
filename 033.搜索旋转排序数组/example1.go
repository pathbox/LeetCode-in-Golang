package LeetCode033

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/solution/yi-wen-jie-jue-4-dao-sou-suo-xuan-zhuan-pai-xu-s-3/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 两种连续递增的情况，判断target是否在连续递增的区间中
		// [left,mid] 连续递增
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1 // 在左侧 [left,mid) 查找
			} else {
				left = mid + 1
			}
		} else { // [mid,right] 连续递增
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1 // 在右侧 (mid,right] 查找
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
