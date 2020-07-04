package LeetCode070

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//方法4 动态规划 + 优化
func climb4(n int) int {
	if n == 1 {
		return 1
	}
	one, two := 1, 2
	for i := 3; i <= n; i++ {
		one, two = two, one+two
	}

	return two
}
