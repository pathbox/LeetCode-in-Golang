package LeetCode394

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" { // [后面跟的是字符串，所以sub存的就是入栈的子字符串
				sub = append(sub, stk[len(stk)-1]) // 得到一个子字符串数组
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ { // 将得到的子字符串数组反转，因为从栈中读取出来的时候子字符串的顺序是反的
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]                      // 子字符串后的栈顶就是[，将其去掉，之后的就是数值
			repTime, _ := strconv.Atoi(stk[len(stk)-1]) // 得到数值
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime) // 将子字符串重复次数得到新的子字符串后，入栈,被之后外层的[]进行处理
			stk = append(stk, t)
		}
	}
	return getString(stk) // 可以用 strings.Join(stk, "")代替
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
