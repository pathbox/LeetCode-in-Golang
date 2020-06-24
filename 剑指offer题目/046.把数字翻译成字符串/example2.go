package Offer046

import "strconv"

func translateNum1(num int) int {
	if num == 0 {
		return 1
	}
	numStr := strconv.Itoa(num)
	pre, cur := 1, 1
	for i := 1; i < len(numStr); i++ {
		tmp := cur
		if numStr[i-1] == '1' || numStr[i-1] == '2' && numStr[i] < '6' {
			cur += pre
		}
		pre = tmp
	}
	return cur
}
