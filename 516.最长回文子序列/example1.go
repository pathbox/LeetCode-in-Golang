package LeetCode516

// https://leetcode-cn.com/problems/longest-palindromic-subsequence/solution/zi-xu-lie-wen-ti-tong-yong-si-lu-zui-chang-hui-wen/

func longestPalindromeSubseq(s string) int {
	if len(s) == 1 {
		return 1
	}

	l := len(s)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	for i :=0; i < l; i++ {
		dp[i][i] = 1
	}

	for i := l-1; i >= 0; i-- {
		j := i+1; j <l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]+2
			} else {
				dp[i][j] = max(dp[i+1][j],dp[i][j-1])
			}
		}
	}

	return dp[0][l-1]
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}