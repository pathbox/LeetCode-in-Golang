package LeetCode394

// 很好理解 数字一个栈，字母一个栈

func decodeString(s string) string {
	alphaStack := make([]byte, 0)
	countStack := make([]int, 0)
	alphaIndex := make([]int, 0)
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			count = count*10 + int(s[i]-'0')
		} else if s[i] == '[' { // [ 后面肯定是字母
			// 可以把当前的数字放入次数数组
			countStack = append(countStack, count)
			count = 0

			alphaIndex = append(alphaIndex, len(alphaStack))
		} else if s[i] == ']' { // 对当前字母段出站
			// 获取当前字母段的次数 然后次数出栈
			c := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			// 获取当前字母段开始的索引 然后索引出栈
			index := alphaIndex[len(alphaIndex)-1]
			alphaIndex = alphaIndex[:len(alphaIndex)-1] // 此时len(alphaIndex)-1之前的字母是已经排好的

			//获取当前的字母段
			str := string(alphaStack[index:])
			// 字母数组出站
			alphaStack = alphaStack[:index] // 会把字母出栈，要不会多计算一次

			for j := 0; j < c; j++ {
				alphaStack = append(alphaStack, []byte(str)...)
			}
		} else {
			alphaStack = append(alphaStack, s[i])
		}
	}
	return string(alphaStack)
}
