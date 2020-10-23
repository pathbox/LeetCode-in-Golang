package LeetCode0812

// 八皇后回溯法

var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = make([][]string, 0)
	queens := make([]int, n) // 存的是每一行queen的列的索引位置
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := make(map[int]bool)                                    // 每一列
	diagonals1, diagonals2 := make(map[int]bool), make(map[int]bool) // 两条对角线
	backtrack(queens, n, 0, columns, diagonals1, diagonals2)

	return solutions
}

func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board) // 得到一种解决方案
		return
	}

	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2)
		// 回溯
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)

	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
