package LeetCode064

func minPathSum(grid [][]int) int {
	l := len(grid)
	if l < 1 {
		return 0
	}
	dp := make([][]int, l)
	for i, arr := range grid {
		dp[i] = make([]int, len(arr))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < l; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 { // 处理两个边缘情况
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	return dp[l-1][len(dp[l-1])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
