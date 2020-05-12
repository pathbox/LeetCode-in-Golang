package LeetCode062

// dp 动态规划
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1 // 初始情况
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j] // 到达dp[i][j]的可能数量等于到达 dp[i][j-1] + dp[i-1][j]的数量，两种情况的和
			}
		}
	}
	return dp[m-1][n-1] // 到达dp[m][n]的可能数量等于到达 dp[m-1][n-1]的数量
}
