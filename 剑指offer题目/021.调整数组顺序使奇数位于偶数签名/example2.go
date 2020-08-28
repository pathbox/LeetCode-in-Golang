package Offer021

func exchange(nums []int) []int {
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j]%2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++ // 发生交换了 i++
		}
		j++
	}
	return nums
}
