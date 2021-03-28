package LeetCode079

func exist(board [][]byte, word string) bool {
	n := len(board)    // 矩阵长度
	m := len(board[0]) // 矩阵宽度

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if backtrace(board, word, i, j, 0) {
				return true // 找到了一个满足情况即可返回true
			}
		}
	}
	return false
}

// k 下一个字符索引
func backtrace(board [][]byte, word string, i, j, k int) bool {
	// 找到了就返回不用继续
	if k == len(word) {
		return true
	}
	// 越界了就说明没有
	if i < 0 || j < 0 || i == len(board) || j == len(board[i]) {
		return false
	}

	if board[i][j] != word[k] { // 字符比较相等
		return false
	}

	tmp := board[i][j]
	// 重置它是为了回溯往回找的时候避免重复使用，干脆，如果找到它是对的，就直接把它置为 空
	// 等结束了之后再重置回来
	board[i][j] = byte(' ') // 避免重复使用

	// 开始上下左右探测
	// left
	if backtrace(board, word, i-1, j, k+1) {
		return true
	}
	// right
	if backtrace(board, word, i+1, j, k+1) {
		return true
	}
	// up
	if backtrace(board, word, i, j-1, k+1) {
		return true
	}
	// down
	if backtrace(board, word, i, j+1, k+1) {
		return true
	}
	board[i][j] = tmp //这次调用没有，将值还原
	return false

}

/*
O(2^n)
分析过程：每个单词的字母在矩阵中都有四个方向查找，因此时间复杂度 4^n ~= O(2^n)

在进入DFS前，设置temp变量记录board[i][j]原来的值，然后把board[i][j]的值原地改为' '
等从DFS出来后，再把board[i][j]改回原来的值
这样就可以避免重复进入
并且因为是原地修改，避免了辅助变量，开销较小

*/
