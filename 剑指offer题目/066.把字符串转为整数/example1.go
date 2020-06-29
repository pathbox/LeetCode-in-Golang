package Offer066

func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	bts := []byte(str)
	var ch byte = ' '
	i := 0
	ch = bts[i]
	for ch == ' ' && i < len(bts)-1 {
		i++
		ch = bts[i]
	}
	bts = bts[i:] // 剔除前面出空

	if len(bts) == 0 || (bts[0] >= 'a' && bts[0] <= 'z') || (bts[0] >= 'A' && bts[0] <= 'Z') {
		return 0
	}
	na := false
	if bts[0] == '0' {
		na = true
		bts = bts[1:]
	} else if bts[0] == '+' {
		bts = bts[1:]
	}

	n := 0
	for _, ch := range bts {
		ch -= '0' // 转为一个整数
		if ch > 9 {
			break
		}
		if n > (1<<31)-1 {
			if na {
				return -2147483648
			} else {
				return (1 << 31) - 1
			}
		}
		n = n*10 + int(ch)
	}
	if n > (1<<31)-1 {
		if na {
			return -2147483648
		} else {
			return (1 << 31) - 1
		}
	}
	if na {
		n = -n
	}
	return n
}
