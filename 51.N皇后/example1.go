package LeetCode51

func solveNQueens(n int) [][]string {
	var tmp = make([][]rune, n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			tmp[r] = append(tmp[r], '.')
		}
	}

	var res [][]string
	backTrack(0, tmp, &res)
	return res
}

func backTrack(row int, tmp [][]rune, res *[][]string) {
	if row == len(tmp) {
		var subRes []string
		for i := 0; i < row; i++ {
			subRes = append(subRes, string(tmp[i]))
		}
		*res = append(*res, subRes)
		return
	}

	for c := 0; c < len(tmp); c++ {
		if !valid(row, c, tmp) {
			continue
		}
		tmp[row][c] = 'Q'
		backTrack(row+1, tmp, res)
		tmp[row][c] = '.'
	}
}

func valid(row, col int, res  [][]rune)bool {
	for i := 0; i < col; i++ {
		if res[row][i] == 'Q' {
			return false
		}
	}

	for i := 0; i < row; i++ {
		if res[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1,col-1; i>=0 && j>= 0; i,j = i-1,j-1 {
		if res[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row -1; col+1; i >=0 &&j <= len(res); i,j = i-1,j+1 {
		if res[i][j] == 'Q' {
			return false
		}
	}
	return true
}