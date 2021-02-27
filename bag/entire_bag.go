package bag

func backpack(nums [][]int, total int) int {
	dp := make([]int, total+1)

	for i := 0; i < len(nums); i++ {
		for j := nums[i][0]; j <= total; j++ {
			dp[j] = max(dp[j], dp[j-nums[i][0]]+nums[i][1])
		}
	}
	return dp[total]
}
