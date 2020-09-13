package LeetCode014

func longestCommonPrefix(strs []string) string {
	strNumber := len(strs)

	if strNumber == 0 {
		return ""
	}

	//取得最短的字符串长度,即最大可能的公共前缀长度
	maxPreLength := len(strs[0])
	for x := 1; x < strNumber; x++ {
		strLength := len(strs[x])
		if strLength < maxPreLength {
			if strLength == 0 {
				//存在空字符串
				return ""
			}
			maxPreLength = strLength // 取最短的item元素长度
		}
	}

	i := 0
I:
	for ; i < maxPreLength; i++ { // 在maxPreLength 下 比较每个元素之间的最长公共前缀 i是字符索引 j是元素索引
		for j := 1; j < strNumber; j++ { // 从i为0的开始，顺着下去两两比较i索引字符是否相等
			if strs[j][i] != strs[j-1][i] { // 如果不相等则退出比较，此时i就是最长的公共前缀的索引
				break I
				// return strs[0][0:i] // 返回最长公共前缀
			}
		}
	}
	return strs[0][0:i] // 返回最长公共前缀
}
