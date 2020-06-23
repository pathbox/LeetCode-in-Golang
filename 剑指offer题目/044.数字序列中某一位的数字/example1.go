package Offer044

import "strconv"

// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/zhe-shi-yi-dao-shu-xue-ti-ge-zhao-gui-lu-by-z1m/

func findNthDigit(n int) int {
	l, c, i := 1, 9, 1
	for n > l*c {
		n -= l * c
		l++
		c *= 10
		i *= 10
	}
	i += (n - 1) / l
	s := []byte(strconv.Itoa(i))
	return int(s[(n-1)%l] - '0')
}
