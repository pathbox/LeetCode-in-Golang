package LeetCode242

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 由于哈希表起始长度为0，所以需要截住 "" == "" 的特殊情况
	if len(s) == 0 {
		return true
	}

	// “哈希表”
	words := make(map[rune]int) // rune在go中是int32别名，任何UTF-8编码（变长）的字符都可以直接用它表示（定长）

	for i := 0; i < len(s); i++ {
		words[rune(s[i])]++
		words[rune(t[i])]--
	}

	// 检查哈希表的最终状态
	for _, v := range words { // 不用检查key，只要发现有值非0，那么就false
		if v != 0 {
			return false
		}
	}

	return true
}
