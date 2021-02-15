package LeetCode494

func findTargetSumWays(nums []int, S int) int {
	numsLen := len(nums)
	var dp = make([]ways, numsLen+1)
	dp[0] = ways{0: 1}
	for i, v := range nums {
		dp[i+1] = ways{}
		for val, times := range dp[i] {
			dp[i+1].add(val+v, times)
			dp[i+1].add(val-v, times)
		}
	}
	return dp[numsLen][S]
}

type ways map[int]int

func (w ways) add(val, times int) {
	if _, ok := w[val]; ok {
		w[val] += times
	} else {
		w[val] = times
	}
}
