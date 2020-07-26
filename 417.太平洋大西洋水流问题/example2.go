package LeetCode417

func pacificAtlantic(matrix [][]int) [][]int {
	res := make([][]int, 0)
	if nil == matrix || len(matrix) == 0 {
		return res
	}

	row, col, dirs := len(matrix), len(matrix[0]), [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	// 是否在区域中
	inArea := func(x, y int) bool {
		return x >= 0 && x < row && y >= 0 && y < col
	}

	var dfs func(x, y int, matrix [][]int, travelled [][]int)
	dfs = func(x, y int, matrix [][]int, travelled [][]int) {
		if travelled[x][y] == 1 {
			return
		}
		travelled[x][y] = 1
		// dirs是四个方向
		for _, d := range dirs {
			// 根据四个方向 得到下一个点
			newx := x + d[0]
			newy := y + d[1]
			// 不在区域 下一个点的高度小于前一个点 已经访问过，都不满足条件，继续下一个循环
			if !inArea(newx, newy) || matrix[x][y] > matrix[newx][newy] || travelled[newx][newy] == 1 {
				continue
			}
			dfs(newx, newy, matrix, travelled)
		}
	}

	pcf, atl := make([][]int, row), make([][]int, row) // 大西洋 太平洋 初始化
	for i := 0; i < row; i++ {
		pcf[i] = make([]int, col)
		atl[i] = make([]int, col)
	}
	// 从外层边缘出发
	for i := 0; i < row; i++ {
		dfs(i, 0, matrix, pcf)
		dfs(i, col-1, matrix, atl)
	}

	for i := 0; i < col; i++ {
		dfs(0, i, matrix, pcf)
		dfs(row-1, i, matrix, atl)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if pcf[i][j] == 1 && atl[i][j] == 1 { // 需要两个都同时满足
				res = append(res, [][]int{{i, j}}...)
			}
		}
	}
	return res
}
