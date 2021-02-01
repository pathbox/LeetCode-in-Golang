package LeetCode392

var (
	dirs          = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	rows, columns int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, columns)
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(anx, dfs(matrix, i, j, memo))
		}
	}
	return ans
}

func dfs(matrix [][]int, row, col int, memo [][]int) int {
	if memo[row][col] != 0 {
		return memo[row][col]
	}
	for _, dir := range dirs {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < columns {
			memo[row][col] = max(memo[row][col], dfs(matrix, newRow, newCol, memo)+1)
		}
	}
	return memo[row][col]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
