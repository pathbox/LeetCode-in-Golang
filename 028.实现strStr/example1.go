package LeetCode028

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	l2 := len(needle)
	for i := 0; i < len(haystack)-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}
