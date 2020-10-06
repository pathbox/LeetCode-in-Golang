package LeetCode322

// https://leetcode-cn.com/problems/coin-change/solution/322-ling-qian-dui-huan-by-leetcode-solution/

// 理解一下i-coin:当前取了coin这个银币出来，dp[i]表示i这么多钱能兑换的银币数量。dp[i] = dp[i-coin]+1 +1表示取出了一个硬币coin
//动态规划（压缩空间）
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i < coin || dp[i-coin] == -1 { // i-coin < 0 过滤
				continue
			}
			count := dp[i-coin] + 1
			if dp[i] > count || dp[i] == -1 { // 取更小的值
				dp[i] = count
			}
		}
	}
	return dp[amount]
}
