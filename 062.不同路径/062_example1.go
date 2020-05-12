package LeetCode062

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	var res []int
	// 初始化,第一行永远是1，第一列也永远是1，因为复用res，所以res[0]永远表示第一列，永远不变
	for i := 0; i < n; i++ {
		res = append(res, 1)
	}
	// 时间复杂度mxn不变
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[j] += res[j-1]
		}
	}
	return res[n-1]
}
