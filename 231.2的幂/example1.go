package LeetCode231

// 如何获取二进制中最右边的 1：x & (-x)。
// 如何将二进制中最右边的 1 设置为 0：x & (x - 1)

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
