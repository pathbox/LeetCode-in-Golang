package LeetCode392

var cache map[[2]int]int = make(map[[2]int]int)

func longestIncreasingPath(matrix [][]int) int {
	if 0 == len(matrix) {
		return 0
	}
	m, n, res := len(matrix), len(matrix[0]), 1
	cache = make(map[[2]int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans := dfs(matrix, i, j)
			res = max(res, ans)
		}
	}
	return res
}

func dfs(matrix [][]int, i int, j int) int {
	if v, ok := cache[[2]int{i, j}]; ok {
		return v
	}
	left, top, right, bottom := 1, 1, 1, 1 // 相当于每回加1

	if j-1 >= 0 && matrix[i][j-1] > matrix[i][j] {
		left += dfs(matrix, i, j-1)
	}
	if i-1 >= 0 && matrix[i-1][j] > matrix[i][j] {
		top += dfs(matrix, i-1, j)
	}
	if j+1 < len(matrix[0]) && matrix[i][j+1] > matrix[i][j] {
		right += dfs(matrix, i, j+1)
	}
	if i+1 < len(matrix) && matrix[i+1][j] > matrix[i][j] {
		bottom += dfs(matrix, i+1, j)
	}

	// ans为上下左右值的最大值
	ans := max(left, max(top, max(right, bottom)))
	cache[[2]int{i, j}] = ans
	return ans
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
