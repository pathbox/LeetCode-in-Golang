package LeetCode1027

// https://leetcode-cn.com/problems/longest-arithmetic-subsequence/solution/bao-li-fa-dong-tai-gui-hua-you-si-lu-miao-shu-by-y/
func longestArithSeqLength(A []int) int {
	length := len(A)
	if length < 3 {
		return 0
	}

	dp := make([]map[int]int, length+1)
	ans := 0
	for i := 0; i < len(dp); i++ {
		dp[i] = make(map[int]int)
	}
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			dis := A[i] - A[j]
			if _, ok := dp[j][dis]; ok {
				dp[i][dis] = dp[j][dis] + 1
			} else {
				dp[j][dis] = 1
				dp[i][dis] = dp[j][dis] + 1
			}
			ans = max(ans, dp[i][dis])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
