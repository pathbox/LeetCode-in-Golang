package LeetCode062

// dfs 递归方法 容易超时
func uniquePaths(m int, n int) int {
	cnt := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x == (m-1) || y == (n-1) {
			cnt++
			return
		} else {
			dfs(x+1, y)
			dfs(x, y+1)
		}
	}
	dfs(0, 0)
	return cnt
}
