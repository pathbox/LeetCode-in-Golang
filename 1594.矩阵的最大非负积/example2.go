package LeetCode1594

func maxProductPath(grid [][]int) int {
	dp := make([][][]int, len(grid))
	for k := range dp {
		dp[k] = make([][]int, len(grid[0]))
	}

	// 先将边缘的情况初始化
	dp[0][0] = []int{grid[0][0], grid[0][0]}
	for i := 1; i < len(dp); i++ {
		if grid[i][0] < 0 {
			dp[i][0] = []int{dp[i-1][0][1] * grid[i][0], dp[i-1][0][0] * grid[i][0]}
		} else {
			dp[i][0] = []int{dp[i-1][0][0] * grid[i][0], dp[i-1][0][1] * grid[i][0]}
		}
	}
	for i := 1; i < len(dp[0]); i++ {
		if grid[0][i] < 0 {
			dp[0][i] = []int{dp[0][i-1][1] * grid[0][i], dp[0][i-1][0] * grid[0][i]}
		} else {
			dp[0][i] = []int{dp[0][i-1][0] * grid[0][i], dp[0][i-1][1] * grid[0][i]}
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if grid[i][j] < 0 {
				dp[i][j] = []int{max(dp[i-1][j][1]*grid[i][j], dp[i][j-1][1]*grid[i][j]), min(dp[i-1][j][0]*grid[i][j], dp[i][j-1][0]*grid[i][j])}
			} else {
				dp[i][j] = []int{max(dp[i-1][j][0]*grid[i][j], dp[i][j-1][0]*grid[i][j]), min(dp[i-1][j][1]*grid[i][j], dp[i][j-1][1]*grid[i][j])}
			}
		}
	}
	if dp[len(grid)-1][len(grid[0])-1][0] < 0 {
		return -1
	}
	return dp[len(grid)-1][len(grid[0])-1][0] % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
