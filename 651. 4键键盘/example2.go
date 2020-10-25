package LeetCode651

func maxA(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1 // 单纯的只是按A键
		// 全选&复制 dp[j-2]，联系复制i-j次 j为若干个c-v的起点
		for j := 2; j < i; j++ {
			dp[i] = max1(dp[i], dp[j-2]*(i-j+1))
		}
	}
	return dp[n]
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// O(N^2)
