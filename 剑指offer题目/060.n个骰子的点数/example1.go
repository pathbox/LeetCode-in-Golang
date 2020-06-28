package Offer060

import "math"

func twoSum(n int) []float64 {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 6*n+1)
	}
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	for i := 1; i <= n; i++ {
		dp[i][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := i + 1; j <= 6*i; j++ {
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 {
					dp[i][j] += dp[i-1][j-k]
				}
			}
		}
	}

	result := make([]float64, 6*n)
	for i := n; i <= 6*n; i++ {
		result[i-1] = float64(dp[n][i]) / math.Pow(6, float64(n))
	}
	return result[n-1:]
}
