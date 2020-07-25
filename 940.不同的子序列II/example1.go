package LeetCode940

import "strings"

// https://leetcode-cn.com/problems/distinct-subsequences-ii/solution/bu-tong-de-zi-xu-lie-ii-by-leetcode/
func distinctSubseqII(S string) int {
	dp := make([]int, len(S)+1)
	dp[0] = 0

	mod := 1000000000 + 7
	for k := 0; k < len(S); k++ {
		tmp := S[:k]
		index := strings.LastIndex(tmp, string(S[k])) // 现在当前的index
		if index < 0 {                                // 表示当前S[k]和前面的tmp没有重复
			dp[k+1] = (2*dp[k] + 1) % mod
			continue
		} else { // 当前S[k]和前面的tmp有重复
			dp[k+1] = (2*dp[k] - dp[index]) % mod // 减去的是S[k]最近重复字符index的dp[index],因为dp[index]的组合是会重复的
		}
	}
	if dp[len(S)] < 0 {
		dp[len(S)] += mod
	}
	return dp[len(S)] % mod
}
