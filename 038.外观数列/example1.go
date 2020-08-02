package LeetCode038

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	l := len(s)
	result := ""
	prevChar := s[0]
	count := 1
	for i := 1; i < l; i++ {
		char := s[i]
		if char == prevChar {
			count++
		} else {
			result += fmt.Sprintf("%d%c", count, prevChar)
			prevChar = char
			count = 1
		}
	}
	result += fmt.Sprintf("%d%c", count, prevChar)
	return result
}
