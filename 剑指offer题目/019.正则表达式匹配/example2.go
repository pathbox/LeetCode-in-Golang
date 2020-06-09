package Offer019

// 二维dp
// dp[i][j]表示s[0:i]和p[0:j]是否匹配
// dp[0][0] = true （s和p都为空）
// dp[i][0](i!=0) = false （s不为空且p为空，匹配必定失败）
// dp[0][j] s为空且p不为空，如果p[j-1]为*且dp[0][j-2]=true则为true（*将p[j-2]忽略）
// 1.如果当前位相等，即p[j-1] == s[i-1]或者p[j-1] == '.' 则dp[i][j] = dp[i-1][j-1] dp[i][j]是否匹配取决于dp[i-1][j-1]
// 2.如果当前位不等
//  (1)当前位p[j-1]为普通字符，或者只剩一个*，匹配失败，dp[i][j] = false为默认值
//  (2)当前位p[j-1] == '*'，且*之前有其他元素
//     ①p前一位与s当前位相等，即p[j-2] == s[i-1]或者p[j-2] == '.'
//       *可以匹配0个，即dp[i][j] = dp[i][j-2]
//       *可以匹配1个，即dp[i][j] = dp[i-1][j-2]
//       *可以匹配多个，即dp[i][j] = dp[i-1][j] （忽略s当前元素）
//       综上dp[i][j] = dp[i-1][j] || dp[i-1][j-2] || dp[i][j-2]
//     ②p前一位与s当前位不等，只能让*匹配0个，即dp[i][j] = dp[i][j-2]
func isMatch(s string, p string) bool {
	rows, cols := len(s), len(p)
	dp := make([][]bool, rows+1)
	for i := 0; i <= rows; i++ {
		dp[i] = make([]bool, cols+1)
	}
	dp[0][0] = true // 都为空是匹配的
	for j := 1; j <= cols; j++ {
		if p[j-1] == '*' && dp[0][j-2] {
			dp[0][j] = true
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				if p[j-1] == '*' && j >= 2 {
					if p[j-2] == s[i-1] || p[j-2] == '.' {
						dp[i][j] = dp[i-1][j] || dp[i-1][j-2] || dp[i][j-2]
					} else {
						dp[i][j] = dp[i][j-2]
					}
				}
			}
		}
	}
	return dp[rows][cols]
}
