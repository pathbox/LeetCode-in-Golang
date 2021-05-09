package LeetCode340

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	n := len(s)
	if n <= 0 || k <= 0 {
		return 0
	}
	m := make(map[byte]int)
	count, res := 0, 1

	for left, right := 0, 0; right < n; right++ {
		if m[s[right]] == 0 {
			count++
		}
		m[s[right]]++

		for count > k {
			m[s[left]]--
			if m[s[left]] == 0 {
				count--
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
