package LeetCode796

import "strings"

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if strings.Contains(A+A, B) {
		return true
	}
	return false
}

func rotateString(A string, B string) bool {
	if A == B {
		return true
	}
	if len(A) != len(B) {
		return false
	}
	for i := 0; i < len(B); i++ {
		B = string(B[1:]) + string(B[0]) // B在一个一个旋转
		if A == B {
			return true
		}
	}
	return false
}
