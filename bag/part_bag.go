package bag

func backpack(nums [][]int, total int) int {
	// nums: [][]int{[]int{5,8,2},[]int{4,3,2}}
	// nums[i][0]表示物品重量 nums[i][1]表示物品价值 nums[i][2]表示最大数量

	dp := make([]int, total+1)

	count := make([]int, total+1)

	for i := 0; i < len(nums); i++ {
		// 每次重置count 每次循环表示一个物品存放的个数，所以要重置
		count = make([]int, total+1)
		for j := nums[i][0]; j <= total; j++ {
			// count[j-nums[i][0]]表示已经放了物品i的个数，如果大于物品i的总个数，跳过
			if count[j-nums[i][0]] < nums[i][2] {
				dp[j] = max(dp[j], dp[j-nums[i][0]]+nums[i][1])

				// 如果放入物品i价值更大，则count[j]需要在countj-nums[i][0](之前放入物品的格式)的基础上加1
				if dp[j] == (dp[j-nums[i][0]] + nums[i][1]) {
					count[j] = count[j-nums[i][0]] + 1
				}
			}
			if dp[j] == 0 || dp[j] < dp[j-1] {
				dp[j] = dp[j-1]
				count[j] = count[j-1]
			}
		}
	}
	return dp[total]
}
