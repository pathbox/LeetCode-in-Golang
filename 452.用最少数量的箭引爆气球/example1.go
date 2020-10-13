package LeetCode452

import "sort"

// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/solution/yong-zui-shao-shu-liang-de-jian-yin-bao-qi-qiu-b-2/
func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return n
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	cnt := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if points[i][0] > points[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		cnt = max(cnt, dp[i])
	}
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
