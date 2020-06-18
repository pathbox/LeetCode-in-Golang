package Offer042

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, sum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] < nums[i] {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
