// https://leetcode-cn.com/problems/sudoku-solver/solution/zi-cong-wo-xue-hui-liao-hui-su-suan-fa-zhong-yu-hu/
//
package LeetCode037

func solveSudoku(board [][]byte) {
	dfs(board, 0, 0)
}

func dfs(board [][]byte, row, col int) bool {
	if row == 9 { // 行遍历完
		return true
	}
	if col == 9 { // 列遍历完从下一行开始
		return dfs(board, row+1, 0)
	}

	if board[row][col] != '.' { // 当前不是空格, 选择下一列
		return dfs(board, row, col+1)
	}

	var k byte
	for k = '1'; k <= '9'; k++ { // 填充1-9的数
		if isValidate(board, row, col, k) { // row col k 这个数是否在当前格子可以填入
			board[row][col] = k         // 当前是空格，用k值进行填充 保护现场
			if dfs(board, row, col+1) { // 遍历下一列
				return true
			}
			board[row][col] = '.' // 回溯
		}
	}
	return false
}

func isValidate(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == c { // 同列是否有相同的数
			return false
		}
		if board[row][i] == c { // 同行是否有相同的数
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}
	return true
}
