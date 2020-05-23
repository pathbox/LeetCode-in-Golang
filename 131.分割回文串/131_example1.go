package LeetCode131

func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}
	ps := make([][]string, 0)
	var subps [][]string
	for i := range s {
		if isPalindrome(s[0 : i+1]) {
			if i == len(s)-1 {
				ps = append(ps, []string{s[0 : i+1]})
			} else {
				subps = partition(s[i+1:])
				for j := range subps {
					ps = append(ps, append([]string{s[0 : i+1]}, subps[j]...)) // 前部分和后部分
				}
			}
		}
	}
	return ps
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] { // 回文字符串的判断
			return false
		}
	}
	return true
}
