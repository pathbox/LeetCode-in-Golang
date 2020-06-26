package Offer049

/*
1. 右窗口不断滑动
2. 窗口内出现了重复的字母，左窗口滑动
3. 统计最长条件
     1. 出现了重复字母
     2. 最后窗口滑动结束，统计一词
*/

func lengthOfLongestSubstring(s string) int {
	var bytes = []byte(s)
	var left, right = 0, 0
	var maxDistince = 0
	for left <= right {
		var m = make(map[byte]bool)
		for k := left; k < right; k++ {
			m[bytes[k]] = true
		}
		for right < len(bytes) {
			if _, ok := m[bytes[right]]; ok {
				if right-left > maxDistince {
					maxDistince = right - left
				}
				left++
				break
			} else {
				m[bytes[right]] = true
				right++
			}
		}
		if right-left > maxDistince {
			maxDistince = right - left
		}
		if right >= len(bytes)-1 {
			break
		}
	}
	return maxDistince
}
