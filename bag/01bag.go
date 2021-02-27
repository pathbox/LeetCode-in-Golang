package bag

func backpack(nums [][]int, total int) int {
	// nums[i][0]表示物品重量 nums[i][1]表示物品价值
	// 有多少物品就有多少行

	// dp[i][j] 表示从下标为[0-i]的物品里任意取，放进容量为j的背包，价值总和最大是多少
	// i: 放i个物品
	// j: 背包的容量
	// total: 指定的背包最大容量
	dp := make([][]int, len(nums))

	// 列代表背包的容量
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, total+1)
	}

	// 第一行放第一个物品，由于里面还没有物品，从能放下这个物品的重量开始，所以最优价值都是第一个物品本身的价值
	for i := nums[0][0]; i < total; i++ {
		dp[0][i] = nums[0][1] // 无论背包的容量多少，初始都只放第一个物品
	}

	// for j := total; j >= nums[0][0]; j-- {
	// 	dp[0][j] = dp[0][j-nums[0][0]] + nums[0][1]
	// }

	// 遍历每个物品,从第二个物品开始
	for i := 1; i < len(nums); i++ { // 遍历物品
		// for j := 0; j <= total; j++ { // 遍历背包容量
		for j := nums[i][0]; j < total; j++ { // 遍历物品重量
			// dp[i-1][j]表示没有放这个物品时的最大价值，就是前一个物品的最大价值。nums[i][1]+dp[i-1][j-nums[i][0]]:第i个物品的价值 + (背包重量-放该i物品之前容许的最大重量的最大价值)
			// if j < nums[i][0] {
			// 	dp[i][j] = dp[i-1][j]
			// } else {
			// 	dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i][0]]+nums[i][0])
			// }

			dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i][0]]+nums[i][1])
		}
	}
	return dp[len(nums)-1][total]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
