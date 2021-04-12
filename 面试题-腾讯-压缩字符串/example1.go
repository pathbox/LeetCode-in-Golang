package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "HG[3|B[2|CA]]F" // HGBCACABCACABCACAF
	// HGBCACABCACABCACAF

	stack := make([]string, 0)
	ptr := 0

	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stack = append(stack, digits)
		} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') || cur == '[' || cur == '|' {
			stack = append(stack, string(cur))
			ptr++
		} else {
			ptr++
			sub := make([]string, 0)
			for stack[len(stack)-1] != "|" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-1-i] = sub[len(sub)-1-i], sub[i]
			}
			stack = stack[:len(stack)-1] // 把 | 出栈
			reqTime, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]                        // 把数字和[出栈
			t := strings.Repeat(strings.Join(sub, ""), reqTime) // 将新构造的字符串入栈
			stack = append(stack, t)
		}
	}
	res := strings.Join(stack, "")
	fmt.Println(res)
	fmt.Println(res == "HGBCACABCACABCACAF")
}

func getDigits(s string, ptr *int) string {
	res := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		res += string(s[*ptr])
	}
	return res
}
