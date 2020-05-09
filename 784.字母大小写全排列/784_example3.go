package LeetCode784

import "strings"

func letterCasePermutation(S string) []string {
	var count uint
	ans := []string{}
	for i := 0; i < len(S); i++ {
		if isLetter(S[i]) {
			count++ //count为字母个数
		}
	}
	for j := 0; j < (1 << count); j++ { //结果中一共有1 << count个字符串 2的count次方
		var c uint
		var tmp string
		for _, v := range S {
			if isLetter(byte(v)) {
				if (j>>c)&1 == 1 { //按位判断该连接大写还是小写字母  掩码的方法 大写小写 0 1
					tmp += strings.ToLower(string(v))
				} else {
					tmp += strings.ToUpper(string(v))
				}
				c++
			} else {
				tmp += string(v)
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

func isLetter(c byte) bool { //判断是否为字母的函数，因为题目说了S仅由数字和字母组成，所以只要大于‘9’就是字母了
	return c > '9'
}
