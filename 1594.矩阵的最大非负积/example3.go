package LeetCode1594

func maxProductPath(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	maxgt := make([][]int, m)
	minlt := make([][]int, m)
	for i := 0; i < m; i++ {
		maxgt[i] = make([]int, n)
		minlt[i] = make([]int, n)
	}

	maxgt[0][0] = grid[0][0]
	minlt[0][0] = grid[0][0]

	for i := 1; i < n; i++ {
		maxgt[0][i] = maxgt[0][i-1] * grid[0][i]
		minlt[0][i] = maxgt[0][i-1] * grid[0][i]
	}

	for i := 1; i < m; i++ {
		maxgt[i][0] = maxgt[i-1][0] * grid[i][0]
		minlt[i][0] = maxgt[i-1][0] * grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] > 0 {
				maxgt[i][j] = max(maxgt[i-1][j], maxgt[i][j-1]) * grid[i][j]
				minlt[i][j] = min(minlt[i-1][j], minlt[i][j-1]) * grid[i][j]
			} else {
				maxgt[i][j] = min(minlt[i-1][j], minlt[i][j-1]) * grid[i][j]
				minlt[i][j] = max(maxgt[i-1][j], maxgt[i][j-1]) * grid[i][j]
			}
		}
	}
	if maxgt[m-1][n-1] < 0 {
		return -1
	}
	return maxgt[m-1][n-1] % 1000000007
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
