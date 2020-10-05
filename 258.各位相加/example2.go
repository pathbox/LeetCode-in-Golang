package LeetCode258

func addDigits(num int) int {
	for num >= 10 {
		num = num%10 + num/10
	}
	return num
}

func addDigits(num int) int {
	for num >= 10 {
		res := 0
		for num > 0 {
			res += num % 10
			num = num / 10
		}
		num = res
	}
	return num
}
