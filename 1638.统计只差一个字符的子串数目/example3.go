package LeetCode1638

func countSubstrings(s, t string) int {
	ans := 0
	m := len(s)
	n := len(t)

	dp := make([][]int, m) // dp[i][j]表示以s[i]和t[j]结尾的所有子串对中，满足恰好只有一个字符不同的字符串对的数目
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	fp := make([][]int, m) // f[i][j]表示以s[i]和t[j]为结尾的子串，最多有多少个连续相同的字符。
	for i := 0; i < m; i++ {
		fp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == t[j] {
				fp[i][j] = 1
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1]
					fp[i][j] += fp[i-1][j-1]
				}
			} else {
				dp[i][j] = 1
				if i > 0 && j > 0 {
					dp[i][j] += fp[i-1][j-1]
				}
			}
			ans += dp[i][j]
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/count-substrings-that-differ-by-one-character/solution/dong-tai-gui-hua-zui-hou-qiu-dpshu-zu-suo-you-yuan/
