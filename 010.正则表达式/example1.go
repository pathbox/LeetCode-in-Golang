package LeetCode010

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}

	if len(p) > 1 && p[1] == '*' {
		switch p[0] {
		case '.':
			for i := 0; i <= len(s); i++ {
				if isMatch(s[i:], p[2:]) {
					return true
				}
			}
		default:
			if isMatch(s, p[2:]) {
				return true
			}

			for i := 0; i < len(s) && s[i] == p[0]; i++ {
				if isMatch(s[i+1:], p[2:]) {
					return true
				}
			}
		}

	} else {
		if s == "" || p == "" {
			return false
		}

		if s[0] == p[0] || p[0] == '.' {
			return isMatch(s[1:], p[1:])
		}
	}
	return false
}
