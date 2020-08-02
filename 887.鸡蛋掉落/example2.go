package LeetCode887

// Mk(T,N)扔的次数 = max{M(k,N-1),M(T-k,N)}+1 T层楼N个鸡蛋,第一个鸡蛋扔到k层时，一共需要确定的次数
// M(T,N) = min{M1 M2 M3... MT}
func superEggDrop(K int, N int) int {
	d := make([][]int, N+1) // n层楼
	for i := range d {
		d[i] = make([]int, K+1) // k个鸡蛋
	}
	for i := 1; i <= N; i++ {
		// 楼层数为i，鸡蛋数量为1时，需要从低到高遍历，共i次
		d[i][1] = i
		for j := 2; j <= K; j++ { // 遍历k个鸡蛋
			d[i][j] = 1 << 32 // 取最小值，先反向赋值一个很大的值

			// 二分法在区间[1,i]里确定一个最优值
			left, right := 1, i
			for left < right {
				// 找dp[k-1][j-1] <= dp[i-mid][j] 的最大值k
				mid := left + (right-left+1)/2
				breakCount := d[mid-1][j-1]
				notBreakCount := d[i-mid][j]
				if breakCount > notBreakCount {
					// 排除法（减治思想）写对二分见第 35 题，先想什么时候不是解
					// 严格大于的时候一定不是解，此时 mid 一定不是解
					// 下一轮搜索区间是 [left, mid - 1]
					right = mid - 1
				} else {
					// 这个区间一定是上一个区间的反面，即 [mid, right]
					// 注意这个时候取中间数要上取整，int mid = left + (right - left + 1) / 2;
					left = mid
				}
			}
			// left在这个下标就是最优的k值，把它代入转移方程即可
			d[i][j] = max(d[left-1][j-1], d[i-left][j]) + 1
		}
	}
	return d[N][K]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 作者：custer-go
// 链接：https://leetcode-cn.com/problems/super-egg-drop/solution/custerxue-xi-bi-ji-dong-tai-gui-hua-by-custer-go-2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	for j := 1; j <= N; j++ {
		for i := 1; i <= K; i++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
			if dp[i][j] >= N {
				return j
			}
		}
	}
	return N
}
