package LeetCode394

import "strings"

var (
	src string
	ptr int
)

func decodeString(s string) string {
	src = s
	ptr = 0
	return getString()
}

func getString() string {
	if ptr == len(src) || src[ptr] == ']' {
		return ""
	}
	cur := src[ptr]
	repTime := 1
	ret := ""
	if cur >= '0' && cur <= '9' {
		repTime = getDigits()
		ptr++
		str := getString()
		ptr++
		ret = strings.Repeat(str, repTime)
	} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
		ret = string(cur)
		ptr++
	}
	return ret + getString()
}

func getDigits() int {
	ret := 0
	for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {
		ret = ret*10 + int(src[ptr]-'0')
	}
	return ret
}
