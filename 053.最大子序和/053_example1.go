package LeetCode053

// nice nums[i]要么以当前值开始，要么等于当前值加上前面序列之和
// 这种概念叫做：前缀和
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] { // nums[i-1]不是原来数组的元素值， 而是0:i-1 最大的累加值
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func maxSubArrayDP(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

func maxSubArray(nums []int) int {
	l := len(nums)
	max := nums[0] // max初始化为nums[0]
	for i := 1; i < l; i++ {
		if nums[i-1] > 0 { // 可以进行累加
			nums[i] += nums[i-1] // 如果nums[i-1]>0 无论nums[i]的值是大于0还是小于0，与其相加都是递增的。nums[i]存的就是到这个i索引之前最大的子序列之和
		} else if nums[i-1] < 0 {
			// 什么也不做
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
