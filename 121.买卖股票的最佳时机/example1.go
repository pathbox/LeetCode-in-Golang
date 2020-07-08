package LeetCode121

import "math"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-can-kao-dian-zan/

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// 创建二维切片
	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, 2)
	}

	// 特殊处理
	dp[0][0] = 0
	dp[0][1] = -float64(prices[0]) // 第0天就买了股票,利润就是负数，因为花去了成本

	for i := 1; i < n; i++ {
		dp[i][0] = math.Max(dp[i-1][0], dp[i-1][1]+float64(prices[i]))
		dp[i][1] = math.Max(dp[i-1][1], -float64(prices[i]))
	}
	return int(dp[n-1][0])
}

/*
dp[i][1] 表示第 i 天 有股票时候的利润；dp[i][0] 表示 第 i 天没有股票时候的利润
dp[i][0] = math.Max(dp[i-1][0], dp[i-1][1] + prices[i])
// Max里的第一项表示第 i - 1 天没有股票
// Max里的第二项表示第 i - 1 天有股票，但是被卖了，所以要加上股票价格，也就是利润
dp[i][1] = math.Max(dp[i - 1][1], -float64(prices[i]))
// Max里的第一项表示第 i - 1 天有股票
// Max里的第二项表示第 i 天买了股票，所以要减去股票价格
*/
