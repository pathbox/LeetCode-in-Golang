package LeetCode276

/*
由于每个栅栏都有两个状态：
0. 与上一个颜色相同 1. 与上一个颜色不同

那么我们可以开辟一个 dp[n][2] 大小的数组来表示:

dp[i][0] // 状态 0 表示与上一个颜色相同

dp[i][1] // 状态 1 表示与上一个颜色不同

状态转移方程也很容易写出来：

dp[i][0] = dp[i - 1][1];

dp[i][1] = (dp[i - 1][0] + dp[i - 1][1]) * (k - 1);

最结果为dp[n - 1][0] + dp[n - 1][1]
*/

func numWays(n int, k int) int {
	if n <= 0 || k <= 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = k

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][1]
		dp[i][1] = (dp[i-1][0] + dp[i-1][1]) * (k - 1)
	}
	return dp[n-1][0] + dp[n-1][1]
}
