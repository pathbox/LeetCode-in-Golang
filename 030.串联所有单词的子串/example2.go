package LeetCode030

func findSubstring(s string, words []string) []int {
	if len(words) < 1 {
		return []int{}
	}

	slen := len(s)
	wlen := len(words)
	k := len(words[0])
	if slen < k*wlen {
		return []int{}
	}

	var res []int
	var mp = make(map[string]int)
	for _, w := range words {
		mp[w]++
	}

	// 移动窗口减少重复检查单词，按单词长度取不同批次
	for i := 0; i < k; i++ {
		var count int
		countermp := make(map[string]int)

		for l, r := i, i; r < slen-k; r = r + k {
			word := s[r : r+k]
			if num, found := mp[word]; found {
				// 如果计数器中单词数目超标，左移指针直至符合数目要求
				for countermp[word] >= num {
					countermp[s[l:l+k]]--
					count--
					l += k
				}
				countermp[word]++
				count++
			} else {
				// 如果当前单词不在词典里，左移指针至下一个单词，左移过程中清理计数
				for l < r {
					countermp[s[l:l+k]]--
					count--
					l += k
				}
				l += k
			}
			if count == wlen {
				res = append(res, l)
			}
		}
	}
	return res
}
