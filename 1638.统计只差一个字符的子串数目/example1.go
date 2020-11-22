package LeetCode1638

func countSubstrings(s, t string) int {
	ans := 0
	m := len(s)
	n := len(t)

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == t[j] {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				dp[i][j] = 1
				// 向后遍历到0
				l, r := i-1, j-1
				for l >= 0 && r >= 0 {
					if s[l] == t[r] {
						dp[i][j] += 1
						l--
						r--
					} else {
						// 如果此时两个字符不相等，跳出循环 因为此时有两个字符不相同了 不需要再循环比较
						break
					}
				}
			}
			ans += dp[i][j]
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/count-substrings-that-differ-by-one-character/solution/dong-tai-gui-hua-zui-hou-qiu-dpshu-zu-suo-you-yuan/
