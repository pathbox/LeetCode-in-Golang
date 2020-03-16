package LeetCode009

import "strconv"

func isPalindrome(x int) bool {
	if x<0 {
		return false
	}
	xArray := []byte(strconv.Itoa(x))
	xLen := len(xArray)
	xHalfLen := xLen/2
	for i := 0; i < xHalfLen; i++{
		if xArray[i]!= xArray[xLen-i-1]{
			return false
		}
	}
	return true
}
