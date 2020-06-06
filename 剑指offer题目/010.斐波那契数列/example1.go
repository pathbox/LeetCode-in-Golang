package Offer010

// 递归方法但是会超时
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
