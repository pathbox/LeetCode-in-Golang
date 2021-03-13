package LeetCode030

func findSubstring(s string, words []string) []int {
	var result []int // 最终结果
	// 特判
	if len(s) == 0 || len(words) == 0 {
		return result
	}
	// 统计words中每个单词出现的次数
	need := make(map[string]int)

	for _, word := range words {
		need[word]++
	}
	wordLen, wordNum, ls := len(words[0]), len(words), len(s)
	totalLen := wordLen * wordNum // 子串总长度

	// 只要循环 wordLen就可以
	for start := 0; start < wordLen; start++ {
		// 窗口左边，窗口右边，词频匹配的单词数
		left, right, match := start, start, 0
		// 窗口中每个单词出现的次数
		window := make(map[string]int)
		for right+wordLen <= ls {
			rightWord := s[right : right+wordLen] // 当前新加入的单词
			right += wordLen                      // 每次移动wordLen
			if need[rightWord] > 0 {
				window[rightWord]++
				// 比较窗口中的单词和need的窗口
				if window[rightWord] == need[rightWord] { // 当前词频相等
					match++
				}
			}

			// 进行判断
			// 如果满足了长度，判断是否满足磁盘
			if right-left == totalLen {
				if match == len(need) {
					result = append(result, left)
				}
				leftWord := s[left : left+wordLen] // 当前左指针移动移除单词
				left += wordLen
				if need[leftWord] > 0 {
					if window[leftWord] == need[leftWord] {
						match--
					}
					window[leftWord]--
				}
			}
		}
	}
	return result
}
