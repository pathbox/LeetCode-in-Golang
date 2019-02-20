package LeetCode010

// "*" 不会出现在p的首位
// "**" 不会出现，但会出现 ".*."" , ".*.." , ".*.*"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}
	// 构造了一个二维数组(矩阵)

	dp[0][0] = true // 表示 s 和 p是空串时，匹配成功

	for i := 2; i < len(dp); i += 2 { // 专门对 *条件匹配，* 条件前面一定会有一个字符，所以用的索引和增长的索引值是2，而不是1。如果i的前面是*，i位设为true
		if p[i-1] == '*' {
			dp[i][0] = true
		} else {
			break
		}
	}

	for i := 1; i < len(dp); i++ {
		if i < len(p) && p[i] == '*' {
			continue
		}

		for j := 1; j < len(dp[0]); j++ {

			if p[i-1] == '*' {
				if p[i-2] == '.' {
					dp[i][j] = dp[i-2][j-1] || dp[i][j-1] || dp[i-2][j]
				} else {
					dp[i][j] = (dp[i-2][j]) || (p[i-2] == s[j-1] && (dp[i-2][j-1] || dp[i][j-1]))
				}
			} else if p[i-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] && p[i-1] == s[j-1]
			}
		}
	}
	return dp[len(p)][len(s)]
}
