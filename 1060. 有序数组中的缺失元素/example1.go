package LeetCode1060

func missingElement(nums []int, k int) int {

	n := len(nums)

	leftK := nums[n-1] - nums[0] - (n - 1)
	if k > leftK { // 说明最后一个元素缺失的不够k
		return nums[n-1] + k - leftK
	} else {
		left, right := 0, n-1
		for left <= right {
			mid := left + (right-left)<<1
			if nums[mid]-nums[0]-mid >= k {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return nums[left-1] + (k - (nums[left-1] - nums[0] - left + 1))
	}
}
