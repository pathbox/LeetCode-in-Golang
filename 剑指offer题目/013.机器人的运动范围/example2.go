package Offer013

func movingCount(m int, n int, k int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	return dfs(m, n, 0, 0, k, dp)
}

func dfs(m, n, i, j, k int, dp [][]int) int {
	if i < 0 || j < 0 || i >= m || j >= n || dp[i][j] == 1 || (sumPos(i)+sumPos(j)) > k {
		return 0
	}

	dp[i][j] = 1 // 当前格子标记为走过

	// 以当前格子上下左右出发寻找
	sum := 1
	sum += dfs(m, n, i, j+1, k, dp)
	sum += dfs(m, m, i, j-1, k, dp)
	sum += dfs(m, m, i+1, j, k, dp)
	sum += dfs(m, m, i-1, j, k, dp)
	return sum
}

func sumPos(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
