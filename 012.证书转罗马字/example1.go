package LeetCode012

func intToRoman(num int) string {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000} // 隐含条件 按照顺序排序
	s := ""
	x := len(intSlice) - 1
	for num != 0 {
		n := num
		for ; x >= 0; x-- {
			if n >= intSlice[x] {
				s += roman[x]
				num -= intSlice[x]
				break
			}
		}
	}
	return s
}

// 罗马字符是先选择能够的最大的字符拼接表示
