package LeetCode052

func totalNQueens(n int) (ans int) {
	columns := make([]bool, n)
	diagonals1 := make([]bool, 2*n-1) // 左上到右下是否有皇后
	diagonals2 := make([]bool, 2*n-1) // 右上到左下是否有皇后
	backtrack := func(row int) {
		if row == n {
			ans++
		}
		for col, hasQueue := range columns {
			d1, d2 := row-col+n-1, row+col
			if hasQueue || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[col] = true
			diagonals1[d1] = true
			diagonals2[d2] = true
			backTrack(row + 1)
			columns[col] = false
			diagonals1[d1] = false
			diagonals2[d2] = false
		}
	}
	backtrack(0)
	return
}

var count int

func totalNQueens(n int) int {
	count = 0
	dfs(n, 0, 0, 0, 0)
	return count
}

func dfs(n, row, col, pie, na int) {
	if row >= n {
		count++
		return
	}

	bits := (^(col | pie | na)) & ((1 << uint(n)) - 1) //获取所有可填空位
	for bits != 0 {
		bit := bits & -bits                               //取位置排在最后的一个空位
		dfs(n, row+1, col|bit, (pie|bit)<<1, (na|bit)>>1) //递归遍历下一行
		bits = bits & (bits - 1)                          //打掉最后位置上的1， 因为该位置已被占用
	}
}

func totalNQueens(n int) int {
	res := 0
	// 生成一个n x n的数组
	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		item := make([]rune, n)
		for k, _ := range item {
			item[k] = '.'
		}
		board[i] = item
	}

	backtrack(board, &res, 0)
	return res
}

func backtrack(board [][]rune, res *int, row int) {
	// 退出条件
	if row == len(board) {
		*res++
		return
	}

	// N叉树遍历
	for col := 0; col < len(board); col++ {
		// 排除不合法的数据
		if isOk(board, row, col) {
			// 做选择
			board[row][col] = 'Q'
			// 进入下一层
			backtrack(board, res, row+1)
			// 撤销选择
			board[row][col] = '.'
		}
	}
}

func isOk(board [][]rune, row, col int) bool {

	// 检查上方数据
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检查左上方数据
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	// 检查右上方
	for i, j := row-1, col+1; i >= 0 && j < len(board); {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}
