package Offer020

import "strings"

var (
	blank  = 0 // 空格
	digit1 = 1 // 数字(0-9) 无前缀
	sign1  = 2 // +/- 无e前缀
	point  = 4 // '.'
	digit2 = 5 // 数字(0-9) 有符号前缀
	e      = 6 // 'e'
	sign2  = 7 // +/- 有e前缀
	digit3 = 8 // 数字(0-9) 有e前缀
)

/* 有限状态自动机
空格 「 」、数字「 0—90—9 」 、正负号 「 +-+− 」 、小数点 「 .. 」 、幂符号 「 ee 」 。

状态定义：

按照字符串从左到右的顺序，定义以下 9 种状态。

开始的空格
幂符号前的正负号
小数点前的数字
小数点、小数点后的数字
当小数点前为空格时，小数点、小数点后的数字
幂符号
幂符号后的正负号
幂符号后的数字
结尾的空格

https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/solution/mian-shi-ti-20-biao-shi-shu-zhi-de-zi-fu-chuan-y-2/
*/
func isNumber(s string) bool {
	s = strings.TrimRight(s, " ")
	dfa := [][]int{
		[]int{blank, digit1, sign1, point, -1},
		[]int{-1, digit1, -1, digit2, e},
		[]int{-1, digit1, -1, point, -1},
		[]int{-1, digit2, -1, -1, e},
		[]int{-1, digit2, -1, -1, -1},
		[]int{-1, digit2, -1, -1, e},
		[]int{-1, digit3, sign2, -1, -1},
		[]int{-1, digit3, -1, -1, -1},
		[]int{-1, digit3, -1, -1, -1},
	}

	state := 0
	for i := 0; i < len(s); i++ {
		var newState int
		switch s[i] {
		case ' ':
			newState = 0
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newState = 1
		case '+', '-':
			newState = 2
		case '.':
			newState = 3
		case 'e':
			newState = 4
		default:
			return false
		}
		state = dfa[state][newState]
		if state == -1 {
			return false
		}
	}
	return state == digit1 || state == digit2 || state == digit3
}
