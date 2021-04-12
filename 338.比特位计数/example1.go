package LeetCode338

// dp 除以2，比如5，除以2应该是2，余1
func countBits(num int) []int {
	dp := make([]int, num+1)
	dp[0] = 0
	if num == 0 {
		return dp
	}
	dp[1] = 1
	for i := 2; i <= num; i++ {
		dp[i] = dp[i/2] + dp[i%2]
	}
	return dp
}

func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}
