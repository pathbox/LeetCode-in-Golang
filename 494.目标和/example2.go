package LeetCode494

func findTargetSumWays(nums []int, S int) int {
	// dp[i][s] = dp[i-1][s-nums[i]] + dp[i-1][s+nums[i]]
	var sum, t int
	length := len(nums)
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	// 注意：绝对值范围超过了sum的绝对值范围则无法得到
	if abs(S) > abs(sum) {
		return 0
	}
	t = 2*sum + 1
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, t)
	}
	dp[0][sum+nums[0]]++
	dp[0][sum-nums[0]]++
	for i := 1; i < length; i++ {
		for j := 0; j < t; j++ {
			// 这里j对应 动态方程里面的 s-nums[i], s+nums[i]
			// 为了避免 -sum的出现，j的范围是 [0, 2*sum]
			var l, r int
			// 注意边界要 =0
			if j-nums[i] >= 0 {
				l = dp[i-1][j-nums[i]]
			}
			if j+nums[i] < t {
				r = dp[i-1][j+nums[i]]
			}
			dp[i][j] = l + r
		}
	}
	return dp[length-1][sum+S]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
