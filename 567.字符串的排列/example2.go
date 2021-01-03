package LeetCode567

func checkInclusion(s1, s2 string) bool {
	window := make(map[byte]int)
	need := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right := 0, 0
	valid := 0

	for right < len(s2) {
		tempAdd := s2[right]
		window[tempAdd]++
		right++

		if _, ok := need[tempAdd]; ok {
			if window[tempAdd] == need[tempAdd] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			tempDel := s2[left]
			left++
			if _, ok := need[tempAdd]; ok {
				if window[tempDel] == need[tempDel] {
					valid--
				}
				window[tempDel]--
			}
		}
	}
	return false
}
