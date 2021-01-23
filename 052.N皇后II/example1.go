package LeetCode052

func totalNQueens(n int) int {
	res := 0 // 在LeetCode中不能用全局变量 会通不过
	var tmp = make([][]rune, n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			tmp[r] = append(tmp[r], '.')
		}
	}

	backTrack(0, &res, tmp)
	return res
}

func backTrack(row int, res *int, tmp [][]rune) {
	if row == len(tmp) {
		*res++
		return
	}

	for c := 0; c < len(tmp); c++ {
		if !valid(row, c, tmp) {
			continue
		}
		tmp[row][c] = 'Q'
		backTrack(row+1, res, tmp)
		tmp[row][c] = '.'
	}
}

func valid(row, col int, res [][]rune) bool {
	for i := 0; i < col; i++ {
		if res[row][i] == 'Q' {
			return false
		}
	}
	for j := 0; j < row; j++ {
		if res[j][col] == 'Q' {
			return false
		}
	}
	// 3.右下角至左上角检测 以(row,col)为端点，检测左上部分，右下部分不用,因为现在只遍历到row行，下面的row都还没有填Q，因为下一行会继续往上检测
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if res[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(res); i, j = i-1, j+1 {
		if res[i][j] == 'Q' {
			return false
		}
	}
	return true
}
