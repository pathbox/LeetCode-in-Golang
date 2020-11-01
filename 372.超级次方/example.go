package LeetCode372

var base = 1337

func superPow(a int, b []int) int {
	tag := a % base
	res := 1
	for i := len(b) - 1; i >= 0; i-- {
		dec := b[i]
		if dec != 0 {
			res = (res * myPow(tag, dec)) % base
		}
		tag = myPow(tag, 10)
	}
	return res
}

func myPow(a, k int) int {
	if k == 0 {
		return 1
	}

	a %= base
	if k%2 == 1 {
		return a * myPow(a, k-1) % base
	} else {
		sub := myPow(a, k/2)
		return sub * sub % base
	}
}
