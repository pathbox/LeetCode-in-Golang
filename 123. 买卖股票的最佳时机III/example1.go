package LeetCode123

func maxProfit(prices []int) int {
	pLen := len(prices)
	if pLen == 0 {
		return 0
	}

	result := 0
	profit := make([][3][2]int, pLen)

	/**
	 *三维数组
	 *profit[ii][kk][jj]
	 *ii 第 ii 天， kk 股票操作了几次 ， jj 是否有股票
	 *最多可以完成 两笔 交易： kk 可以为 0 1 2 次操作 ， jj可以为 0 ，1
	 **/
	/**
	*第一天 初始化
	*第 1 天 操作 k 次 没有股票，所以初始值为 0
	*第 1 天 操作i 次 有股票， 所以初始值为 - prices[0]
	注意，最后一个dp[i][2][1] 实际是不存在的，因为交易两次后，就不能再买入了
	https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/wu-chong-shi-xian-xiang-xi-tu-jie-123mai-mai-gu-pi/
	**/
	profit[0][0][0], profit[0][0][1] = 0, -prices[0]
	profit[0][1][0], profit[0][1][1] = 0, -prices[0]
	profit[0][2][0], profit[0][2][1] = 0, -prices[0]

	for i := 1; i < pLen; i++ {
		//第 i 天 0 次交易 没有股票最大利润 = 第 i-1 天 0 次交易 没有股票最大利润
		profit[i][0][0] = profit[i-1][0][0]
		// 第一次买入：从初始状态转换而来，或者第一次买入后保持不动
		profit[i][0][1] = max(profit[i-1][0][1], profit[i-1][0][0]-prices[i])

		// 第一次卖出：从第一次买入转换而来，或者第一次卖出后保持不动
		profit[i][1][0] = max(profit[i-1][1][0], profit[i-1][0][1]+prices[i])
		// 第二次买入：从第一次卖出转换而来，或者第二次买入后保持不动
		profit[i][1][1] = max(profit[i-1][1][1], profit[i-1][1][0]-prices[i])

		//# 第 i 天 2 次交易 无股票最大利润 = max(第 i-1 天 2次交易 无股票最大利润 , 第 i-1 天 1 次交易 有股票最大利润 + 当天股票价格prices[i]（卖出）)
		profit[i][2][0] = max(profit[i-1][2][0], profit[i-1][1][1]+prices[i])

	}

	result = max(profit[pLen-1][0][0], max(profit[pLen-1][1][0], profit[pLen-1][2][0]))
	return result
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
