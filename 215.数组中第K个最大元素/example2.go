package LeetCode215

func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums) - 1

	for {
		if left >= right { // 重要
			return nums[right]
		}
		p := partition(nums, left, right)
		if p+1 == k {
			return nums[p]
		} else if p+1 < k {
			left = p + 1
		} else {
			right = p - 1
		}
	}
}

func partition(nums []int, left int, right int) int {
	pivot := nums[right]
	for i := left; i < right; i++ {
		if nums[i] > pivot {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	return left
}
