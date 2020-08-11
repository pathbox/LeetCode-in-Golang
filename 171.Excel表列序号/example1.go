package LeetCode171

func titleToNumber(s string) int {
	res := 0
	runes := []rune(s)
	for _, c := range runes {
		res = 26*res + (int(c-'A') + 1)
	}
	return res
}
