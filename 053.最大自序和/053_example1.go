package LeetCode053

func maxSubArray(nums []int) int {
	l := len(nums)
	max := nums[0] // max初始化为nums[0]
	for i := 1; i < l; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1] // 如果nums[i-1]>0 无论nums[i]的值是dayu0还是小于0，与其相加都是递增的。nums[i]存的就是到这个i索引之前最大的子序列之和
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
