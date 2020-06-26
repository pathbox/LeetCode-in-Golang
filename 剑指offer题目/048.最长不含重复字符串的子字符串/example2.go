package Offer048

/*
1. 右窗口不断滑动
2. 窗口内出现了重复的字母，左窗口滑动
3. 统计最长条件
     1. 出现了重复字母
     2. 最后窗口滑动结束，统计一词
*/

func lengthOfLongestSubstring(s string) int {
	// q store index of first
	q := []int{}
	m := make(map[byte]int)
	maxLength := 0
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && v > 0 {
			for len(q) != 0 && s[q[0]] != s[i] { // 不断的出队列，直到重复的元素不在队列中
				m[s[q[0]]]--
				q = q[1:]
			}
			m[s[q[0]]]--
			q = q[1:]
		}
		m[s[i]]++
		q = append(q, i)
		if len(q) > maxLength {
			maxLength = len(q)
		}
	}
	return maxLength
}
