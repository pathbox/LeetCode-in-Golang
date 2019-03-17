pacakge LeetCode060

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	// 存放结果字符,最后只输出一种排列结果，所以只要一个数组即可，不需要map或二维数组
	res := make([]byte, n)

	rec := make([]byte,n)
	for i := 0; i < n; i++{
		rec[i] = byte(i)+'1' // 构造这个序列数数组
	}

	// 由于排列的序号是从 1 开始的。
	// k 需要减去 1 ，好从 0 开始表示
	k--

	base := 1
	for i := 2; i < n; i++{
		base *= i
	}


	for i := 0; i < n-1; i++{
		idx := k / base // 每一位在rec中的idx索引.为什么用k / base, 最高位可取1,2,3,4中的一个，每个数字出现3！= 6次,从竖轴方向看数字，k / base 相当于竖方向跳过了几个数，得到的数-1(数组从0索引开始)，就是k对应序列数此时的索引

		res[i]= rec[idx] // 取出该位的数放入res
		// 从 rec 中去除已经使用的数 rec[idx]
		rec = append(rec[:idx],rec[idx+1:]...)

		k %= base
		base /= (n-i-1)
	}
	// 不要忘记最后一个数
	res[n-1] = rec[0] // rec只剩下最后一个数

	return string(res)
}