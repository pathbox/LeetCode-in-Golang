package Offer048

// dp
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s))

	dp[0] = 1

	memo := make(map[byte]int)
	max := 1
	memo[s[0]] = 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			dp[i] = 1
		} else {
			lastAppear := -1
			if idx, ok := memo[s[i]]; ok {
				lastAppear = idx
			}
			if i-lastAppear > dp[i-1] {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = i - lastAppear

			}
		}

		memo[s[i]] = i

		if max < dp[i] {
			max = dp[i]
		}
	}
	return max

}
