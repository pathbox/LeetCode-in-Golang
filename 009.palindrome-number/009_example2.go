package LeetCode009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	val := 0
	temp := x
	for x != 0 {
		val = val*10 + x%10
		x = x / 10
	}
	if val == temp {
		return true
	} else {
		return false
	}
}
