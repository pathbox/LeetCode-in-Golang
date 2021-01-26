package LeetCode392

var cache map[[2]int]int = make(map[[2]int]int)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n, res := len(matrix), len(matrix[0]), 1
	cache := make(map[[2]int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans := dfs(matrix, i, j)
			res = max(res, ans)
		}
	}
	return res
}

func dfs(matrix [][]int, i, j int) int {
	if v, ok := cache[[2]int{i, j}]; ok { // 已经访问过了的节点不用再继续
		return v
	}
	left, top, right, bottom := 1, 1, 1, 1
	if j-1 >= 0 && matrix[i][j-1] > matrix[i][j] {
		left += dfs(matrix, i, j-1)
	}
	if i-1 >= 0 && matrix[i-1][j] > matrix[i][j] {
		top += dfs(matrix, i-1, j)
	}
	if j+1 < len(matrix[0]) && matrix[i][j+1] > matrix[i][j] {
		right += dfs(matrix, i, j+1)
	}
	if i+1 < len(matrix) && matrix[i][i+1] > matrix[i][j] {
		bottom += dfs(matrix, i+1, j)
	}
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
