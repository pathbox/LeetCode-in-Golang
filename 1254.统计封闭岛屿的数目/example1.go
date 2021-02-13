package LeetCode1254

func closedIsland(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	res := 0
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if grid[i][j] == 0 {
				if dfs(grid, i, j) {
					res++
				}
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) bool {
	row := len(grid)
	col := len(grid[0])
	if i < 0 || j < 0 || i >= row || j >= col {
		return false
	}
	if grid[i][j] == 1 {
		return true
	}
	grid[i][j] = 1
	up := dfs(grid, i-1, j)
	down := dfs(grid, i+1, j)
	left := dfs(grid, i, j-1)
	right := dfs(grid, i, j+1)
	if up && down && left && right {
		return true
	}
	return false
}
