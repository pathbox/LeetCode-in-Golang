package LeetCode279

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n // 最多是n次， 1+1+1+1+1+...+1 n次
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1) // 动态转移方程
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

// https://leetcode-cn.com/problems/perfect-squares/solution/wan-quan-bei-bao-wen-ti-by-wayne-63/
