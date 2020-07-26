package LeetCode807

func maxIncreaseKeepingSkyline(grid [][]int) int {
	res := 0
	maxR := make([]int, len(grid))
	maxC := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > maxR[i] {
				maxR[i] = grid[i][j]
			}
			if grid[i][j] > maxC[j] {
				maxC[j] = grid[i][j]
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res += min(maxR[i], maxC[j]) - grid[i][j]
		}
	}
	return res
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
