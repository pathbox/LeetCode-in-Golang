package LeetCode052

var res int

func totalNQueens(n int) int {
	var tmp = make([][]rune, n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			tmp[r] = append(tmp[r], '.')
		}
	}

	backTrack(0, tmp)
	return res
}

func backTrack(row int, tmp [][]rune) {
	if row == len(tmp) {
		res++
		return
	}

	for c := 0; c < len(tmp); c++ {
		if !valid(row, c, tmp) {
			continue
		}
		tmp[row][c] = 'Q'
		backTrack(row+1, tmp)
		tmp[row][c] = '.'
	}
}

func valid(row, col int, tmp [][]rune) bool {
	for i := 0; i < len(tmp); i++ {
		if tmp[row][i] == 'Q' {
			return false
		}
	}

	for i := 0; i < len(tmp); i++ {
		if tmp[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if tmp[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j <= len(tmp); i, j = i-1, j+1 {
		if tmp[i][j] == 'Q' {
			return false
		}
	}

	return true
}
