package LeetCode153

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right { // 实际上是不会跳出循环，当 left==right 时直接返回
		if nums[left] <= nums[right] { // 如果 [left,right] 递增，直接返回, nums[left]就是最小
			return nums[left]
		}
		mid := left + (right-left)>>1
		if nums[left] <= nums[mid] { // [left,mid] 连续递增，则在 [mid+1,right] 查找
			left = mid + 1
		} else {
			right = mid // [left,mid] 不连续，在 [left,mid] 查找
		}
	}
	return -1
}
