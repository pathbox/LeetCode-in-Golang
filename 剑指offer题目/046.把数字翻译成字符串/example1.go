package Offer046

func translateNum(num int) int {
	pn := -1    // 前一个数字(0-9)
	cn := 0     // 当前数字
	ans := 1    // 默认1个
	last := ans // 前一位的数量(i-2)
	for num > 0 {
		cn = num % 10
		num = num / 10
		// 可以结合[10-25]
		if cn != 0 && pn >= 0 && cn*10+pn <= 25 {
			ans, last = ans+last, ans
		} else {
			last = ans
		}
		pn = cn
	}
	return ans
}
