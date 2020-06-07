package Offer010

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	var (
		a    = 0
		b    = 1
		temp int
	)
	for i := 2; i <= n; i++ {
		temp = b
		b = (a + b) % 1000000007
		a = temp
	}
	return b
}
