package LeetCode030

// 方法一：双map，从每一个位置开始遍历，重复检查单词
func findSubstring(s string, words []string) []int {
	var res []int
	if len(words) < 1 {
		return res
	}

	wlen := len(words)
	k := len(words[0])
	if len(s) < k*wlen {
		return res
	}

	mp := make(map[string]int)
	for _, w := range words {
		mp[w]++
	}

	for i := 0; i < len(s)-k*wlen+1; i++ { // 每次s移动一个字符
		var count int
		mp2 := make(map[string]int)
		// 遍历每个单词 进行比较
		for multi := 0; multi < wlen; multi++ {
			start := i + multi*k // 每次跳一个单词的位置
			word := s[start : start+k]
			if num, found := mp[word]; found && num > mp2[word] {
				mp2[word]++
				count++
			} else {
				break
			}
		}

		if count == wlen { // 如果每个单词都遍历了，满足条件
			res = append(res, i)
		}
	}
	return res
}
