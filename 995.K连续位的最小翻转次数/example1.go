package LeetCode995

func minKBitFlips(A []int, K int) int {
	var (
		res    int
		flip   int                   // 记录当前窗口目前总共进行了多少次翻转
		window = make([]int, len(A)) // 记录某个位置是否进行了翻转
	)

	for i := 0; i < len(A); i++ {
		// 出窗
		if i >= K {
			flip -= window[i-K]
		}
		// 入窗
		// 下面这个判断包含了两种情况：
		// 如果 flippedTime 是奇数，且 A[i] == 1 就需要翻转
		// 如果 flippedTime 是偶数，且 A[i] == 0 就需要翻转
		if (flip%2)^A[i] == 0 {
			// 如果过了i+k-1还要修改后面窗口，那就不行
			if i+K > len(A) {
				return -1
			}
			flip++
			res++
			window[i] = 1
		}
	}
	return res
}
