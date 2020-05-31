package LeetCode036

func isValidSudoku(board [][]byte) bool {
	var row, col, sbox [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				index_box := (i/3)*3 + j/3
				if row[i][num] == 1 { // i行上是否有num数字
					return false
				} else {
					row[i][num] = 1
				}
				if col[j][num] == 1 {
					return false
				} else {
					col[j][num] = 1
				}
				if sbox[index_box][num] == 1 {
					return false
				} else {
					sbox[index_box][num] = 1
				}
			}
		}
	}
	return true
}
