package LeetCode1638

/*
O(mn*min(mn)) O(1)
*/
func countSubstrings(s, t string) int {
	m := len(s)
	n := len(t)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff := 0
			for k := 0; i+k < m && j+k < n; k++ {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff > 1 {
					break
				}
				if diff == 1 {
					ans++
				}
			}
		}
	}
	return ans
}
