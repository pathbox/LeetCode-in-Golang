package Offer014

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	dp := make([]int, n+1)
	// dp数组第i个元素表示长度为i的绳子剪成若干段或不剪,之后各段长度乘积的最大值
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	max := 0
	for i := 4; i <= n; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			res := dp[j] * dp[i-j]
			if max < res {
				max = res
			}
			dp[i] = max
		}
	}
	max = dp[n]
	dp = nil
	return max
}
