package LeetCode678

func checkValidString(s string) bool {
	if len(s) == 0 {
		return true
	}
	if s[0] == ')' { // 右括号开头的栈
		return false
	}
	left := []int{} // 保存的是下标
	star := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left = append(left, i)
		}
		if s[i] == '*' {
			star = append(star, i)
		}

		if s[i] == ')' {
			ll := len(left)
			ls := len(star)
			if ll > 0 {
				left = left[:ll-1] // 出栈
			} else if ls > 0 {
				star = star[:ls-1] // 当成左括号进行匹配出栈
			} else {
				return false // 表示没有左括号和*与右括号匹配
			}
		}
	}

	if len(left) == 0 { // 如果左括号刚好匹配完了，说明符合，返回true
		return true
	}
	for i := len(left) - 1; i >= 0; i-- { // 遍历left栈，如果出现left中的下标值大于star的栈顶，说明有 *(的情况存在，是不合法配对,或者没有star与其匹配了 此时是把*当成右括号
		if len(star) == 0 || star[len(star)-1] < left[i] {
			return false
		}
		star = star[:len(star)-1] // 匹配掉一个star
	}
	return true
}
