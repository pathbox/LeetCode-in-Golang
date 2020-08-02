package LeetCode887

func superEggDrop(K int, N int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	var (
		dp = make([][]int, K+1)
		m  = 0
	)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]int, N+1)
	}
	for dp[K][m] != N {
		m++
		for k := 1; k < K+1; k++ {
			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
			if dp[k][m] > N {
				return m
			}
		}
	}
	return m
}

// 作者：haohao95
// 链接：https://leetcode-cn.com/problems/super-egg-drop/solution/golang-dp-100-by-haohao95/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
