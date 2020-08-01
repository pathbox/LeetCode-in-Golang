package LeetCode887

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
			// 遍历鸡蛋扔在f层的情况
			for f := 1; f <= i; f++ {
				// 设临界值为 t
				// d[f-1][j-1]为鸡蛋碎了，t在[1,f]这个区间中，共f层，剩余j-1个鸡蛋
				// d[i-k][j]为鸡蛋未碎，t在[k+1,i]这个区间中，共i-k层，剩余j个鸡蛋
				// 在两种情况中取较大值(考虑最坏情况)加上本次，更新d[i][j]
				d[i][j] = min(d[i][j], max(d[f-1][j-1], d[i-f][j])+1)
			}
		}
	}
	return d[N][K]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
