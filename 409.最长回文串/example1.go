package LeetCode409

func longestPalindrome(s string) int {
	hash := make(map[byte]int)
	sum := 0

	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	for _, v := range hash {
		sum += (v / 2) * 2 // 对2整除，得到多少个可以构成回文串的组合字符. 回文字符偶数个性质
	}

	if sum == len(s) {
		return sum
	} else {
		return sum + 1 // 可以+1字符串放到中间，所以+1
	}
}
