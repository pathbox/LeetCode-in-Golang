package LeetCode032

func longestValidParentheses(s string) int {
	max := 0
	dp := make([]int, len(s))
	for i, l := 1, len(s); i < l; i++ {
		if s[i] == ')' {
			if k := i - dp[i] - 1; k >= 0 && s[k] == '(' {
				m := dp[i] + 2 + dp[k]
				if i+1 < l {
					dp[i+1] = m
				}
				if m > max {
					max = m
				}
			}
		}
		// s[i] == '(' dp[i] = 0
	}
	return max
}
