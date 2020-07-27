package LeetCode393

func validUtf8(data []int) bool {
	i, l := 0, 0
	for i < len(data) {
		if data[i] > 128 {
			if data[i]>>5 == 6 {
				l = 1
			} else if data[i]>>4 == 14 {
				l = 2
			} else if data[i]>>3 == 30 {
				l = 3
			} else {
				return false
			}
			if i+l >= len(data) {
				return false
			}
			for k := 0; k < l; k++ {
				i++
				if data[i]>>6 != 2 {
					return false
				}
			}

		}
		i++
	}

	return true
}
