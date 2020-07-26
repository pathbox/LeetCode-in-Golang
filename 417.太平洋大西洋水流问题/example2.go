package LeetCode417

func pacificAtlantic(matrix [][]int) [][]int {

	res := [][]int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	row, col := len(matrix), len(matrix[0])

	// 设置位
	setBit := func(n, i int) int {
		return n | (1 << i)
	}
	// 验证位
	checkBit := func(n, i int) bool {
		return (n&(1<<i) != 0)
	}
	// 太平洋的标记位
	const Pacific = 0
	// 大西洋的标记位
	const Atlantic = 1
	// 标记矩阵，如果 Pacific 位被设置，则此顶点可达太平洋，大西洋同理
	m := make([][]int, row)
	for i := 0; i < row; i++ {
		m[i] = make([]int, col)
	}

	var dfs func(x, y, value, ocean int)
	dfs = func(x, y, value, ocean int) {

		// 递归结束条件一：y、x超过限制
		if x >= row || y >= col {
			return
		}
		if x < 0 || y < 0 {
			return
		}

		// 递归结束条件二：当前顶点已经被访问且设置了标记
		// 此时就无需再从这个顶点进行递归了
		if checkBit(m[x][y], ocean) {
			return
		}

		if ocean == Pacific && (y == 0 || x == 0) {
			// 太平洋的边界
			m[x][y] = setBit(m[x][y], ocean)
		} else if ocean == Atlantic && (y == col-1 || x == row-1) {
			// 大西洋的边界
			m[x][y] = setBit(m[x][y], ocean)
		} else if matriy[x][y] >= value {
			// 当前顶点能到达下一个顶点
			m[x][y] = setBit(m[x][y], ocean)
		} else {
			// 递归结束条件三：当前顶点无法到达下一个顶点
			return
		}

		// 上
		dfs(x-1, y, matriy[x][y], ocean)
		// 左
		dfs(x, y-1, matriy[x][y], ocean)
		// 下
		dfs(x+1, y, matriy[x][y], ocean)
		// 右
		dfs(x, y+1, matriy[x][y], ocean)
	}
	// 查找太平洋
	dfs(0, 0, -1, Pacific)
	// 查找大西洋
	dfs(row-1, col-1, -1, Atlantic)
	// 找出既能到达太平洋，也能到达大西洋的顶点
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if checkBit(m[i][j], Pacific) && checkBit(m[i][j], Atlantic) {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
