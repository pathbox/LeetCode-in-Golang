package LeetCode014

func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	if strsLen == 1 {
		return strs[0]
	}
	return commonPrefixLong(strs, 0, len(strs)-1)
}

func commonPrefixLong(strs []string, left, right int) string {
	if left == right {
		return strs[left]
	}
	middle := (left + right) / 2
	leftWord := commonPrefixLong(strs, left, middle)
	rightWord := commonPrefixLong(strs, middle+1, right)
	return commonPrefix(leftWord, rightWord)
}

func commonPrefix(result string, word string) string {
	if len(result) > len(word) {
		result = result[0:len(word)] // 取长度小的
	}
	if len(result) < 1 {
		return result
	}
	for j := 0; j < len(result); j++ {
		if result[j] != word[j] {
			result = result[0:j]
			break
		}
	}
	return result
}
