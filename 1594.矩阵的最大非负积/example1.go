package LeetCode1594

import "math"

func maxProductPath(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	nr := len(grid)
	nc := len(grid[0])
	var dfs func(row, col, value int)
	var res int = math.MinInt64
	dfs = func(row, col, value int) {
		// 开头进行满足条件的判断
		value *= grid[row][col]
		if value == 0 || row == nr-1 && col == nc-1 {
			if value > res {
				res = value
			}
			return
		}
		// 继续递归
		if col < nc-1 {
			dfs(row, col+1, value)
		}
		if row < nr-1 {
			dfs(row+1, col, value)
		}
	}
	dfs(0, 0, 1)

	if res < 0 {
		return -1
	}
	return res % 1000000007
}
