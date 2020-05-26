package LeetCode187

func findRepeatedDnaSequences(s string) []string {
	st := 0
	ed := 10
	m := make(map[string]int)
	for ed <= len(s) {
		m[s[st:ed]]++
		st++
		ed++
	}
	res := make([]string, 0)
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
