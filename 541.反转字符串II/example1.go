package LeetCode541

func reverseStr(s string, k int) string {
	str := []byte(s)
	size := len(str)
	var i int
	for i < size {
		l := i
		r := i + k - 1
		if r >= size-1 {
			r = size - 1
		}
		for l < r && l < size {
			str[l], str[r] = str[r], str[l]
			l++
			r--
		}
		i = i + 2*k
	}
	return string(str)
}
