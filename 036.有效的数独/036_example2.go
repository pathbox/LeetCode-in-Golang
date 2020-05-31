package LeetCode036

func isValidSudoku(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1') // 当前数字的二进制数位 位置
			bi := i/3 + j/3*3
			if (row[i]&cur)|(col[j]&cur)|(block[bi]&cur) != 0 {
				return false
			}
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}
