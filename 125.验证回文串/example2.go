package LeetCode125

import "unicode"

func isPalindrome(s string) bool {
	// 判断是否是需要比较的合法回文字符,
	// 题目说空格是合法的回文字符很误导人, 其实这个地方需要排除空格
	isValid := func(v rune) bool {
		return unicode.IsDigit(v) || unicode.IsLetter(v)
	}

	i, j := 0, len(s)-1
	for i < j {
		vi, vj := rune(s[i]), rune(s[j])

		if !isValid(vi) && !isValid(vj) { // 都不符合则移动两个指针
			i++
			j--
		} else if !isValid(vi) { // 前面不符合移动 i
			i++
		} else if !isValid(vj) { // 后面不符合移动 j
			j--
		} else if unicode.ToUpper(vi) != unicode.ToUpper(vj) {
			return false // 如果都合法如果不相等则直接返回
		} else { // 如果相等则移动两个指针
			i++
			j--
		}
	}

	return true
}
