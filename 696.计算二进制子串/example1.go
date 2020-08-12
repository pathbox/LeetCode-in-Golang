package LeetCode696

/*
找到连续两个不同的数字的交界，然后左右两边长度最短的连续相同数字的长度，就是这个边界能组成的子串数量。
那问题就变成了找到所有交界，然后看左右两边谁更短，然后把这个短的数字加到结果即可。
*/
func countBinarySubstrings(s string) int {
	counts := []int{}
	ptr, n := 0, len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		counts = append(counts, count)
	}
	ans := 0
	for i := 1; i < len(counts); i++ {
		ans += min(counts[i], counts[i-1]) // 相邻较小的值
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
