package LeetCode139

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j, v := range wordDict {
			l := len(v) // 每个单词的长度
			if dp[i] == true && (i+1 <= n && s[i:i+l] == wordDict[j]) {
				dp[i+1] = true
			}
		}
	}
	return dp[n]
}
