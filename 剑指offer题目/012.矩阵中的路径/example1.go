package Offer012

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, word string, k int) bool {
	if board[i][j] != word[k] { // 每次递归进来都会先进行比较，比较成功才往下走继续递归,否则直接返回
		return false
	}
	if k == len(word)-1 { // 匹配到了所有word字符 返回true
		return true
	}
	temp := board[i][j]                            // 当前索引位置
	board[i][j] = byte(' ')                        // 表示当前位置不能再重复遍历
	if 0 <= i-1 && dfs(board, i-1, j, word, k+1) { // 向上
		return true
	}
	if i+1 < len(board) && dfs(board, i+1, j, word, k+1) { // 向下
		return true
	}
	if 0 <= j-1 && dfs(board, i, j-1, word, k+1) { // 向左
		return true
	}
	if j+1 < len(board[0]) && dfs(board, i, j+1, word, k+1) { // 向右
		return true
	}
	board[i][j] = temp // 上面的递归都没有得到结果，则回溯
	return false       // 比较都没成功返回false
}
