package LeetCode1249

import "strings"

func minRemoveToMakeValid(s string) string {
	sum := 0
	lenth := len(s)
	tmp := []rune(s)
	for i := 0; i < lenth; i++ {
		if tmp[i] == '(' {
			sum++
		} else if tmp[i] == ')' {
			sum--
			if sum < 0 {
				tmp[i] = '0'
				sum = 0
			}
		}
	}
	sum = 0
	for i := lenth - 1; i >= 0; i-- {
		if tmp[i] == ')' {
			sum++
		} else if tmp[i] == '(' {
			sum--
			if sum < 0 {
				tmp[i] = '0'
				sum = 0
			}
		}
	}
	return strings.ReplaceAll(string(tmp), "0", "")
}

// 链接：https://leetcode-cn.com/problems/minimum-remove-to-make-valid-parentheses/solution/go-bu-shi-yong-zhan-shi-jian-onkong-jian-o1-by-men/
