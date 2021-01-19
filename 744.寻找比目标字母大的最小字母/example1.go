package LeetCode744

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	for l < r {
		mid := l + (r-l)>>1
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if letters[l] > target {
		return letters[l]
	}

	return letters[0] // 如果没有则返回第一个
}
