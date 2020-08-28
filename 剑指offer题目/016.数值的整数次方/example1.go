package Offer016

// 本题是经典的分治法，用递归（recursion）和迭代（iteration）都可以完成。
func help(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n&1 == 0 { // 表示n是偶数可以开平方
		return help(x*x, n/2)
	} else { //  n&1 == 1
		return x * help(x, n-1)
	}
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return help(1.0/x, -n)
	} else if n == 0 {
		return 1
	} else {
		return help(x, n)
	}
}
