package LeetCode1254

func closedIsland(grid [][]int) int {
	num, rowNum, colNum := 0, len(grid), len(grid[0])
	for row, v1 := range grid {
		for col, v2 := range v1 {
			if v2 == 1 {
				continue
			}
			// 如果满足四面环水，则孤岛数量+1
			if dfs(grid, rowNum, colNum, row, col) {
				num++
			}
		}
	}

	return num
}

// 以(i,j)点出发，寻找上下左右周围是否是水,是则返回true,如果是陆地0,则继续移动遍历,如果触碰到边界了,则说明无法构成孤岛
func dfs(grid [][]int, rowNum, colNum, row int, col int) (b bool) {
	// 如果触碰到边界了，无法构成孤岛
	if row < 0 || col < 0 || row >= rowNum || col >= colNum {
		return false
	}
	// 如果碰到水了，返回正确
	if grid[row][col] == 1 {
		return true
	}
	// 经过的地方置为海洋
	grid[row][col] = 1
	// 需要四个方向都环水
	up := dfs(grid, rowNum, colNum, row-1, col)
	down := dfs(grid, rowNum, colNum, row+1, col)
	left := dfs(grid, rowNum, colNum, row, col-1)
	right := dfs(grid, rowNum, colNum, row, col+1)
	return up && down && left && right
}
