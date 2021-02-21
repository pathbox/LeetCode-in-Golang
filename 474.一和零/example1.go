package LeetCode474

func findMaxForm(strs []string, m int, n int) int {
	statistic := func(str string) (int, int) {
		var zero, one = 0, 0
		for _, char := range str {
			if char == '0' {
				zero++
			} else {
				one++
			}
		}
		return zero, one
	}
	length := len(strs)
	dp := make([][][]int, length+1)
	//初始化
	for i := 0; i <= length; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//进行动态规划
	for i := 1; i <= length; i++ {
		//计算0,1数量
		zero, one := statistic(strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				//先将背包当前状态设为直接继承为上一个物品状态
				dp[i][j][k] = dp[i-1][j][k]
				//当满足条件时，进行判断装入还是不装
				if j >= zero && k >= one {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zero][k-one]+1)
				}
			}
		}
	}
	return dp[length][m][n]
}
