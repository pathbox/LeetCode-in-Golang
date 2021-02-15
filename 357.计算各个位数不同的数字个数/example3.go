package LeetCode357

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	dp := make([]int, n+1)
	dp[1] = 10
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 9*math_a(9, i-1)
	}
	return dp[n]
}

func math_a(a, b int) int {
	cnt := 1
	for i := 0; i < b; i++ {
		cnt *= a - i
	}
	return cnt
}
