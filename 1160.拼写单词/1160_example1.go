package LeetCode1160

func countCharacters(words []string, chars string) int {
	if len(words) == 0 {
		return 0
	}

	// 构建字符哈希表
	chs := make([]byte, 'z'-'A'+1)
	for _, ch := range chars {
		chs[ch-'A']++
	}

	res := 0

	for _, str := range words {
		if isContained(str, chs) {
			res += len(str)
		} // 如果包含
	}
	return res
}

func isContained(str string, chs []byte) bool {
	ws := make([]byte, 'z'-'A'+1)
	for _, w := range str {
		ws[w-'A']++
	}
	for i := 0; i < len(ws); i++ {
		if len(ws) > len(chs) {
			return false
		}
		if ws[i] > chs[i] {
			return false
		}
	}
	return true
}

// 解题思路：
// 哈希表；
// 数组的长度
// 注意数组的长度，如果只考虑小写字母就 26 了，考虑大写字母的话，z - A + 1;
//对于一个单词 word，只要其中的每个字母的数量都不大于 chars 中对应的字母的数量，那么就可以用 chars 中的字母拼写出 word。所以我们只需要用一个哈希表存储 chars 中每个字母的数量，再用一个哈希表存储 word 中每个字母的数量，最后将这两个哈希表的键值对逐一进行比较即可
