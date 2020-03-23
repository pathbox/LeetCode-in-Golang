package LeetCode189

func rotate(nums []int, k int) {
	k %= len(nums) // 是个环移动 所以用%取余
	ans := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	nums = append(nums[:0], ans...)
}
