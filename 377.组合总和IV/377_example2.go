package LeetCode377

func combinationSum4(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, val := range nums {
			if i >= val {
				dp[i] = dp[i] + dp[i-val]
			}
		}
	}
	return dp[target]
}
