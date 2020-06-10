package Offer016

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1.0 / x
		n = -n
	} else if n == 0 {
		return 1
	}
	ans := 1.0
	for n > 1 {
		if n&1 == 1 {
			ans *= x
			n--
		} else {
			x = x * x
			n /= 2
		}
	}
	ans *= x
	return ans
}
