package LeetCode130

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rowNum, colNum := len(board), len(board[0])
	var queue [][]int
	// 所有边界元素加入队列
	for row := 0; row < rowNum; row++ {
		if board[row][0] == 'O' {
			queue = append(queue, []int{row, 0})
		}
		if board[row][colNum-1] == 'O' {
			queue = append(queue, []int{row, colNum - 1})
		}
	}
	for col := 1; col < colNum-1; col++ {
		if board[0][col] == 'O' {
			queue = append(queue, []int{0, col})
		}
		if board[rowNum-1][col] == 'O' {
			queue = append(queue, []int{rowNum - 1, col})
		}
	}
	// ⽅向数组 d 是上下左右搜索的常⽤⼿法
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		row, col := head[0], head[1]
		board[row][col] = 'A'
		for _, v := range dir {
			newRow, newCol := row+v[0], col+v[1]
			if newRow < 0 || newCol < 0 || newRow >= rowNum || newCol >= colNum || board[newRow][newCol] != 'O' {
				continue
			}
			queue = append(queue, []int{newRow, newCol})
		}

	}
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			if board[row][col] == 'A' {
				board[row][col] = 'O'
			} else if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}
}
