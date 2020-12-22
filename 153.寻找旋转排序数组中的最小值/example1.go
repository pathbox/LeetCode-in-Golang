package LeetCode153

/*
// 最小值就是拐点
// 1.找到数组的中间元素 mid。
// 2.如果中间元素 > 数组第一个元素，我们需要在 mid 右边搜索变化点。说明左部分是正常顺序，拐点不可能在左部分
// 3.如果中间元素 < 数组第一个元素，我们需要在 mid 左边搜索变化点。说明拐点会在左部分
当我们找到变化点时停止搜索，当以下条件满足任意一个即可：
nums[mid] > nums[mid + 1]，因此 mid+1 是最小值。

nums[mid - 1] > nums[mid]，因此 mid 是最小值
*/

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	if nums[right] > nums[left] {
		return nums[left]
	}

	for left <= right {
		mid := left + (right-left)/2
		// 找到了拐点 正常情况：nums[mid-1]<=nums[mid]<=nums[mid+1]，其他情况就是拐点
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] > nums[0] { // 和第一个nums[0]元素比较，而不是nums[left]元素
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}
