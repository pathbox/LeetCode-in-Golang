package LeetCode091

import "strconv"

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	} else if s[0] == '0' {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1] // 第一次：s[i] 和 dp[i-1]个组合可以组成新的组合
		}
		if s[i-2] == '1' || s[i-2] == '2' && s[i-1] < '7' { // 第二次：s[i]s[i-1] 和 dp[i-2] 可以组成新的组合
			dp[i] = dp[i] + dp[i-2]
		}
	}
	return dp[n]
}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		self, _ := strconv.Atoi(s[i : i+1])
		num, _ := strconv.Atoi(s[i-1 : i+1])
		if self == 0 && num != 10 && num != 20 {
			//非法
			return 0
		} else if self == 0 {
			//只能和i-1一起编
			pre, cur = cur, pre
		} else if num <= 26 && self != num {
			//可以单独编，也可以和i-1一起编
			pre, cur = cur, pre+cur
		} else {
			//只能单独编
			pre = cur
		}
	}
	return cur
}
