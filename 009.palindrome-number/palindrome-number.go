package LeetCode009

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 { // 负数肯定不是
		return false
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
