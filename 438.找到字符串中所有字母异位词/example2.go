package LeetCode0438

func findAnagrams(s string, p string) []int {
	window := make(map[byte]int)

	// 初始化 待匹配数据的数据状态
	need := make(map[byte]int)
	for v := range p {
		need[p[v]]++
	}
	needVal := len(need)

	// 滑动窗口状态
	left, right := 0, 0
	// 窗口与目标字符串对比状态
	valid := 0
	res := make([]int, 0)

	for right < len(s) {
		c := s[right]
		right++
		// 新增的字符在目标结果中
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			// 找到目标更新结果
			if needVal == valid {
				res = append(res, left)
			}
			d := s[left]
			left++
			// 如果删除的字符在目标字符串里面 则变更状态
			if _, ok := need[d]; ok {
				valid--
			}
			// 如果删除的字符在目标字符串里面 则变更状态
			window[d]--
		}
	}
	return res
}
