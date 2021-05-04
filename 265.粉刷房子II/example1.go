package LeetCode265

import "math"

func minCostII(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	n := len(costs)
	K := len(costs[0])

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, K)
	}
	for i := 1; i < n; i++ {
		for j := 0; j < K; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	for j := 0; j < K; j++ {
		dp[0][j] = costs[0][j] // 第1个房子的情况
	}

	for i := 1; i < n; i++ {
		for j := 0; j < K; j++ {
			for k := 0; k < K; k++ {
				if j == k { // 上一个房子的颜色和当前选择的房子颜色不能相等
					continue
				}
				dp[i][j] = min(dp[i-1][k]+costs[i][j], dp[i][j])
			}
		}
	}
	return min(dp[n-1]...)
}

func minCostII(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	m := len(costs[0])
	if m == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := range costs[0] {
		dp[0][i] = costs[0][i]
	}

	for i := 1; i < n; i++ {
		mincolor1, mincolor2 := 0, 0
		if dp[i-1][0] < dp[i-1][1] {
			mincolor1, mincolor2 = 0, 1
		} else {
			mincolor1, mincolor2 = 1, 0
		}
		for k := 2; k < m; k++ {
			if dp[i-1][k] < dp[i-1][mincolor1] {
				mincolor1, mincolor2 = k, mincolor1
			} else if dp[i-1][k] < dp[i-1][mincolor2] {
				mincolor2 = k
			}
		}
		for j := 0; j < m; j++ {
			if j == mincolor1 {
				dp[i][j] = dp[i-1][mincolor2] + costs[i][j]
			} else {
				dp[i][j] = dp[i-1][mincolor1] + costs[i][j]
			}

		}
	}
	return min(dp[n-1]...)
}

func min(lists ...int) int {
	a := lists[0]
	for i := 1; i < len(lists); i++ {
		if lists[i] < a {
			a = lists[i]
		}
	}
	return a
}
