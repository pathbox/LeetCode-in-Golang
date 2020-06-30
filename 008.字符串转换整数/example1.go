package LeetCode008

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	//去掉收尾空格
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' { // 只处理数字，字母不处理
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 { //首字符正负的判断
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else { // 非数字字符，退出
			break
		}
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return sign * result
}
