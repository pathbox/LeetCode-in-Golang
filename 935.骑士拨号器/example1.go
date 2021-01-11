package LeetCode935

func knightDialer(N int) int {
	dp := make([][]int, N)
	dp[0] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 1; i < N; i++ {
		dp[i] = make([]int, 10)
		dp[i][0] = (dp[i-1][4] + dp[i-1][6]) % 1000000007
		dp[i][1] = (dp[i-1][6] + dp[i-1][8]) % 1000000007
		dp[i][2] = (dp[i-1][7] + dp[i-1][9]) % 1000000007
		dp[i][3] = (dp[i-1][4] + dp[i-1][8]) % 1000000007
		dp[i][4] = (dp[i-1][0] + dp[i-1][3] + dp[i-1][9]) % 1000000007
		dp[i][5] = 0
		dp[i][6] = (dp[i-1][0] + dp[i-1][1] + dp[i-1][7]) % 1000000007
		dp[i][7] = (dp[i-1][2] + dp[i-1][6]) % 1000000007
		dp[i][8] = (dp[i-1][1] + dp[i-1][3]) % 1000000007
		dp[i][9] = (dp[i-1][2] + dp[i-1][4]) % 1000000007
	}
	sum := 0
	for i := 0; i < 10; i++ {
		sum = (sum + dp[N-1][i]) % 1000000007
	}
	return sum
}
