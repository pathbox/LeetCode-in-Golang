package LeetCode377

/*
动态规划
dp[nums...] == 1
dp[i] = dp[i] + dp[i - nums...]
*/
func combinationSum4(nums []int, target int) int {
	return DP(nums, target)
}

func DP(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i < target; i++ {
		for _, val := range nums {
			if i == val {
				dp[i] += 1
				continue
			}
			if i > val && dp[i-val] != 0 {
				dp[i] += dp[i-val]
			}
		}
	}
	return dp[target]
}
