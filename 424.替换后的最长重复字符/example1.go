package LeetCode424

func characterReplacement(s string, k int) int {
	if s == "" {
		return 0
	}
	left, right, result, maxCount := 0, 0, 0, 0

	count := make(map[byte]int)

	for right < len(s) {
		count[s[right]]++
		maxCount = max(maxCount, count[s[right]]) // 当前窗口内的最多重复字符的个数,只需要知道窗口中最多重复字符的个数即可
		if right-left+1-maxCount > k {            // 需要替换的字符个数就是当前窗口的大小减去窗口中数量最多的字符的数量
			count[s[left]]-- //缩小窗口
			left++
		} //当窗口内可替换的字符数小于等于k时，我们需要根据该窗口长度来确定是否更新result
		result = max(result, right-left+1) // 每次都计算一次结果，窗口大小只会增不会减
		right++                            // 始终都会往右走
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
