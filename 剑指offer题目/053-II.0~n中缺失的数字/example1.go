package Offer053

// 缺失的数字等于 “右子数组的首位元素” 对应的索引 或是左子数的末尾元素

func missingNumber(nums []int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] != mid {
			//nums是有序数组，如果mid和数字不相同就在左边查找
			right = mid
		} else {
			//如果mid和数字相同，说明左边是连续的有序数组
			//缺失的数字就在右边查找，left向上取整+1
			left = mid + 1
		}
	}
	return left
}
