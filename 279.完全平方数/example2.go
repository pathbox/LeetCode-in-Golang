package LeetCode279

import "math"

func numSquares(n int) int {
	m := int(math.Sqrt(float64(n)))
	dp := make([]int, n+1)

	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
	}

	for i := 1; i <= m; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

// https://leetcode-cn.com/problems/perfect-squares/solution/wan-quan-bei-bao-wen-ti-by-wayne-63/
