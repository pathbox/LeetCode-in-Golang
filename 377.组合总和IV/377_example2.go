package LeetCode377

/*
动态规划算法通常用于求解具有某种最优性质的问题。在这类问题中，可能会有许多可行解。每一个解都对应于一个值，我们希望找到具有最优值的解。动态规划算法与分治法类似，其基本思想也是将待求解问题分解成若干个子问题，先求解子问题，然后从这些子问题的解得到原问题的解。与分治法不同的是，适合于用动态规划求解的问题，经分解得到子问题往往不是互相独立的。对于分治法求解的问题，子问题的相互独立仅仅是同层级的子问题间没有互相依赖。但对于动态规划而言，同层级的子问题可能会依赖相同的低层级问题，这就导致低层级问题可能会被计算多次
*/
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
