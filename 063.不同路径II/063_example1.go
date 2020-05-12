package LeetCode063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid) // 行的数量
	if m < 1 {
		return 0
	}
	n := len(obstacleGrid[0]) // 列的数量
	if n < 1 {
		return 0
	}
	if obstacleGrid[0][0] == 1 { // 初始位置就有障碍物 直接返回0的情况
		return 0
	}
	dp := make([][]int, m) // m行 n列
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 && j != 0 {
				if obstacleGrid[i][j] == 0 {
					dp[i][j] = dp[i][j-1]
				}
			} else if i != 0 && j == 0 {
				if obstacleGrid[i][j] == 0 {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				if obstacleGrid[i][j] == 0 {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}
