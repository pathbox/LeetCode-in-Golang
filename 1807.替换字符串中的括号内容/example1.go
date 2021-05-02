package LeetCode1807

func evaluate(s string, knowledge [][]string) string {
	m := make(map[string]string)
	ret := ""
	for i := 0; i < len(knowledge); i++ {
		m[knowledge[i][0]] = knowledge[i][1]
	}

	sig := false
	temp := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			sig = true
			continue
		} else if s[i] == ')' {
			sig = false
			if _, ok := m[temp]; ok {
				ret += m[temp]
			} else {
				ret += "?"
			}
			temp = ""
			continue
		}
		if sig {
			temp += string(s[i])
		} else {
			ret += string(s[i])
		}
	}
	return ret
}
