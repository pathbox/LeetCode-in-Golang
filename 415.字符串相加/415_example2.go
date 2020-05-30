package LeetCode415

// 两个大数相加
func addStrings(num1 string, num2 string) string {
	var i int = len(num1) - 1
	var j int = len(num2) - 1
	var r byte = 0
	var s = []byte{}
	for i >= 0 || j >= 0 || r != 0 {
		if i >= 0 {
			r += (num1[i] - '0')
			i--
		}
		if j >= 0 {
			r += (num2[j] - '0')
			j--
		}

		s = append(s, ((r % 10) + '0'))
		r /= 10
	}
	var n int = len(s) - 1
	for i := 0; i < n; i++ { // 数组头尾对调，得到想要顺序的排序
		s[i], s[n] = s[n], s[i]
		n-- // 这个不能少，要不然就需要找中间位置进行对调
	}

	return string(s)

}
