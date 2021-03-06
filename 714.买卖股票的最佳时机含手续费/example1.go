package LeetCode714

func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// i天不持有股票：i-1天不持有股票 i天在i-1持有股票的情况下卖了
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		// i天持有股票： i-1天持有股票 i-1天不持有股票 i天购买
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
