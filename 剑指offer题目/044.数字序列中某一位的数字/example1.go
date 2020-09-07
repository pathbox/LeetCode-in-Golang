package Offer044

import "strconv"

// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/zhe-shi-yi-dao-shu-xue-ti-ge-zhao-gui-lu-by-z1m/
// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/golang-shuang-bai-by-eggsy-3-7/

func findNthDigit(n int) int {
	var digst = 1
	var count = 9
	var start = 1
	for n-count > 0 {
		n -= count
		digst += 1
		start *= 10
		count = digst * start * 9
	}
	var num = start + (n-1)/digst
	var res = strconv.Itoa(num)[(n-1)%digst] - '0'
	return int(res)
}
