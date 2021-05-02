package LeetCode1100

func numKLenSubstrNoRepeats(S string, K int) int {
	n := len(S)
	if n < K {
		return 0
	}
	m := make(map[byte]int)
	res := 0
	for left, right := 0, 0; right < n; right++ {
		for m[S[right]] > 0 {
			m[S[left]]--
			left++
		}
		m[S[right]]++
		l := right - left + 1
		if l == K {
			res++
			m[S[left]]--
			left++
		}
	}

	return res
}
