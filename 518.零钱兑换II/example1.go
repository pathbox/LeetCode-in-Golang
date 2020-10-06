package LeetCode518

// https://leetcode-cn.com/problems/coin-change-2/solution/ling-qian-dui-huan-iihe-pa-lou-ti-wen-ti-dao-di-yo/

// 零钱问题是组合问题，爬楼梯问题是排列问题，因为:1+2和2+1的硬币组合现实生活中是一样的意义，而1阶梯+2阶梯  2阶梯+1阶梯，由于顺序不同，现实生活中爬楼梯的情况也是不同的
// 是组合的结果
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i < coin {
				continue
			}
			dp[i] = dp[i] + dp[i-coin]
		}
	}
	return dp[amount]
}

/* 这是排列的结果
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {

			if i < coin {
				continue
			}
			dp[i] = dp[i] + dp[i-coin]
		}
	}
	return dp[amount]
}

*/
