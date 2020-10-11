package LeetCode1143

// https://leetcode-cn.com/problems/longest-common-subsequence/solution/1143-zui-chang-gong-gong-zi-xu-lie-by-ming-zhi-sha/
// dp[i][j]:表示在字符串s[0...i]中和字符串s[0...j]中最长公共子序列的长度为dp[i][j]。
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] { // dp中的i j 整体往后移了一位i j 为0 的索引位置数据是0
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
