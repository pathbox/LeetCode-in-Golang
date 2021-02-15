package LeetCode494

func findTargetSumWays(nums []int, S int) int {
	var (
		count int = 0
		dfs   func(int, int)
	)

	dfs = func(i, sum int) {
		if i == len(nums) {
			if sum == S {
				count++
			}
		} else {
			dfs(i+1, sum-nums[i])
			dfs(i+1, sum+nums[i])
		}
	}
	dfs(0, 0)
	return count
}
