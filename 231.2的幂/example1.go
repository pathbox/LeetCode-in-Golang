package LeetCode231

// https://leetcode-cn.com/problems/power-of-two/solution/power-of-two-er-jin-zhi-ji-jian-by-jyd/
// 如何获取二进制中最右边的 1：x & (-x)。
// 如何将二进制中最右边的 1 设置为 0：x & (x - 1)

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
	// return n > 0 && n&(-n) == n
}
