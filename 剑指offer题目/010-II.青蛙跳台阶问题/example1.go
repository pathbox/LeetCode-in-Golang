package Offer010

// 和斐波那契数列的区别 f[0] = 0 q[0] = 1
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	dp := make([]int, n+1) // 这里要提前设置长度为n+1
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}
