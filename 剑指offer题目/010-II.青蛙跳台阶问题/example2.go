package Offer010

// 用前后指针两个数值也可以
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	a, b, temp := 1, 1, 0
	for i := 2; i <= n; i++ {
		temp = b
		b = (a + b) % 1000000007
		a = temp
	}
	return b
}
