package Offer020

// 动态规划法
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, 0)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-1]) % 1000000007
	}
	return dp[n]
}
