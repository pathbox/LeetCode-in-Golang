package LeetCode089

// n增加时，将上次的结果逆序并在高位补1后追加在后面即可
func grayCode(n int) []int {
	var res []int
	// 默认只有一个0
	res = append(res, 0)
	if n == 0 {
		return res
	}
	if n >= 1 {
		res = append(res, 1)
	}
	base := 1
	for i := 2; i <= n; i++ { //
		// 反转前面的结果，在前面加1
		l := len(res)
		base = base * 2
		for j := l - 1; j >= 0; j-- {
			res = append(res, res[j]+base)
		}
	}
	return res
}
