package LeetCode172

/*
5和任何偶数相乘个位都得0
任何偶数都是2的倍数，并且在阶乘数中2的个数要比5多，因此只要考虑5的个数
所以直接除去得到5的个数
*/

func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n /= 5
	}
	return res
}
