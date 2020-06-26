package Offer047

func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	left, top := 0, 0
	// i, j 是当前位置， top left 是当前位置的上部和左部
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			top, left = 0, 0

			if i-1 >= 0 {
				top = grid[i-1][j]
			}
			if j-1 >= 0 {
				left = grid[i][j-1]
			}

			// 谁更大，累加它
			if top > left {
				grid[i][j] += top
			} else {
				grid[i][j] += left
			}
		}
	}
	return grid[m-1][n-1]
}
