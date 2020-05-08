package LeetCode784

func letterCasePermutation(S string) []string {
	var B uint = 0
	for _, r := range S {
		if isLetter(r) {
			B++
		}
	}
	var ans = make([]string, 0, 1<<B)
	for bits := 0; bits < 1<<B; bits++ {
		var b uint = 0
		runes := make([]rune, len(S))
		i := 0
		for _, r := range S {
			if isLetter(r) {
				if ((bits >> b) & 1) == 1 {
					runes[i] = toLowerCase(r)
				} else {
					runes[i] = toUpperCase(r)
				}
				b++
			} else {
				runes[i] = r
			}
			i++
		}
		ans = append(ans, string(runes))
	}
	return ans
}

func isLetter(r rune) bool {
	return r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z'
}

func toLowerCase(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		r += 32
	}
	return r
}

func toUpperCase(r rune) rune {
	if r >= 'a' && r <= 'z' {
		r -= 32
	}
	return r
}
