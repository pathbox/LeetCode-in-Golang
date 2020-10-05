package LeetCode377

/*
动态规划
dp[nums...] == 1
dp[i] = dp[i] + dp[i - nums...]
此题的状态转移方程为dp[i] = dp[i-nums[0]]+dp[i-nums[1]]+...dp[i-nums[len-1]],条件为i>=nums[j];
/**
     * 这里状态定义就是题目要求的，并不难，状态转移方程要动点脑子，也不难：
     * 状态转移方程：dp[i]= dp[i - nums[0]] + dp[i - nums[1]] + dp[i - nums[2]] + ... （当 [] 里面的数 >= 0）
     * 特别注意：dp[0] = 1，表示，如果那个硬币的面值刚刚好等于需要凑出的价值，这个就成为 1 种组合方案
     * 再举一个具体的例子：nums=[1, 3, 4], target=7;
     * dp[7] = dp[6] + dp[4] + dp[3]
     * 即：7 的组合数可以由三部分组成，1 和 dp[6]，3 和 dp[4], 4 和dp[3];
     *
     * @param nums
     * @param target
     * @return
     */
*/
func combinationSum4(nums []int, target int) int {
	return DP(nums, target)
}

func DP(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
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
