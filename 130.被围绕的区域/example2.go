package LeetCode130

var rowNum, colNum int

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rowNum, colNum = len(board), len(board[0])
	// 对所有边界做处理
	for i := 0; i < rowNum; i++ {
		dfs(board, i, 0)
		dfs(board, i, colNum-1)
	}
	for i := 1; i < colNum-1; i++ {
		dfs(board, 0, i)
		dfs(board, rowNum-1, i)
	}
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			// 被标记过的的说明是和边界相连的
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= rowNum || y < 0 || y >= colNum || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)
}
