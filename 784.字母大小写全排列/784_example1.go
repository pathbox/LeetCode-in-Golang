package LeetCode784

import (
	"strings"
	"unicode"
)

func letterCasePermutation(S string) []string {
	list := []string{""}
	helper(S, &list)
	return list
}

func helper(S string, L *[]string) {
	if S == "" {
		return
	}
	length := len(*L)
	if unicode.IsDigit(rune(S[0])) {
		for i := 0; i < length; i++ {
			(*L)[i] += string(S[0]) // 数字直接加到每个元素的末尾
		}
	} else {
		for i := 0; i < length; i++ {
			*L = append(*L, (*L)[i])                 // 复制操作
			(*L)[i] += strings.ToLower(string(S[0])) // 给前半部分每个元素末尾加上小写字符
		}
		for i := length; i < Len(*L); i++ {
			(*L)[i] += strings.ToUpper(string(S[0])) // 给前半部分每个元素末尾加上大写字符
		}
	}
	if len(S) > 1 {
		helper(S[1:], L)
	} else if len(S) == 1 {
		helper("", L)
	}
	return
}
