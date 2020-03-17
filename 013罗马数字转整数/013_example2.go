package LeetCode013

func getVal(c rune) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	var cs = []rune(s)
	var len = len(cs)
	var pre = getVal(cs[len-1])
	var ret = pre
	for i := len - 2; i >= 0; i-- {
		var n = getVal(cs[i])
		if n >= pre {
			ret += n
		} else {
			ret -= n
		}
		pre = n
	}
	return ret
}
