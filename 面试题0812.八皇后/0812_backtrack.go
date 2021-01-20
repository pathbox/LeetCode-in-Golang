package LeetCode0812

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

// 1. 回溯
func backTrack(row int, tmp [][]rune, res *[][]string) {
	if row == len(tmp) { // 满足一种情况
		var subRes []string
		for i := 0; i < row; i++ {
			subRes = append(subRes, string(tmp[i])) // copy一个结果，不直接使用tmp
		}
		*res = append(*res, subRes)
		return
	}

	for c := 0; c < len(tmp); c++ {
		if !valid(row, c, tmp) {
			continue
		}
		tmp[row][c] = 'Q' // 可以放queens
		backTrack(row+1, tmp, res)
		tmp[row][c] = '.'
	}
}

// 2.校验  都是校验上半部分即可，下半部分还没有Q的安排
func valid(row, col int, res [][]rune) bool {
	for i := 0; i < col; i++ {
		// 1.列检测
		if res[row][i] == 'Q' {
			return false
		}
	}
	for j := 0; j < row; j++ {
		// 2.行检测
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

	// 4.左下角至右上角检测
	for i, j := row-1, col+1; i >= 0 && j < len(res); i, j = i-1, j+1 {
		if res[i][j] == 'Q' {
			return false
		}
	}
	return true
}
