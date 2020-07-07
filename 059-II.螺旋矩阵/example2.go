package LeetCode059

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	fill(result, 1, 0, n)
	return result
}

func fill(r [][]int, start, row, n int) {
	if start > n*n {
		return
	}
	if start == n*n {
		r[row][row] = start
		return
	}
	for i := row; i < n-1-row; i++ {
		r[row][i] = start
		start++
	}
	for i := row; i < n-1-row; i++ {
		r[i][n-1-row] = start
		start++
	}
	for i := n - 1 - row; i > row; i-- {
		r[n-1-row][i] = start
		start++
	}
	for i := n - 1 - row; i > row; i-- {
		r[i][row] = start
		start++
	}

	fill(r, start, row+1, n)

}
