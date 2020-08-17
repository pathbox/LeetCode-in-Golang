package LeetCode416

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}

		}
	}
	return dp[sum]
}
