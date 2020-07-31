package LeetCode1254

func closedIsland(grid [][]int) int {
	num, rowNum, colNum := 0, len(grid), len(grid[0])
	for row, v1 := range grid {
		for col, v2 := range v1 {
			if v2 == 1 { // 水
				continue
			}
			// 如果满足四面环水，则孤岛数量+1
			if dfs(grid, rowNum, colNum, row, col) { // 传入此时的坐标和已知参数
				num++
			}
		}
	}
	return num
}

func dfs(grid [][]int, rowNum, colNum, row, col int) (b bool) {
	// 如果触碰到边界了，无法构成孤岛 先判断边界
	if row < 0 || col < 0 || row >= rowNum || col >= colNum {
		return false
	}

	// 在没有碰到边界的情况下，如果碰到水了，返回正确
	if grid[row][col] == 1 {
		return true
	}
	// 经过的地方置为海洋 避免重复遍历
	grid[row][col] = 1
	// 需要四个方向都环水 四个方向都true
	up := dfs(grid, rowNum, colNum, row-1, col)
	down := dfs(grid, rowNum, colNum, row+1, col)
	left := dfs(grid, rowNum, colNum, row, col-1)
	right := dfs(grid, rowNum, colNum, row, col+1)
	return up && down && left && right
}
