package LeetCode130

// 搜索与边界相通的O，它们不会被X替换
func dfs(board [][]byte, i, j int) {
	// 边界判断:不是边界或者已经搜索过
	// board[i][j] == '#': 说明已经搜索过
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	// 和边界联通的O，先替换成#， 后续换回O
	board[i][j] = '#'
	// 上下左右继续广度优先搜索
	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// 从边缘开始搜索, 判断是否在边上，只从边上开始，如果边上都是X，那么里面的所有O都换成X就行
			// 如果边上有，那么之后就会被替换成#暂存，之后会被恢复成O，意味着不换成X
			isEdge := i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1
			if isEdge && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	// 再替换，这里的O都应该是被包围的
	// # 是和边界联通的，即不被包围的
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}
