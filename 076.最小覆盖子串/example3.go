package LeetCode076

import "math"

func minWindow(s string, t string) string {
	// 保存滑动窗口字符集
	win := make(map[byte]int)
	need := make(map[byte]int)
	left, right := 0, 0
	match := 0
	start, end := 0, 0
	min := math.MaxInt64
	var c byte

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	for right < len(s) {
		// 1.处理right++事件
		c = s[right]
		right++
		if need[c] != 0 {
			win[c]++

			if win[c] == need[c] {
				match++
			}
		}

		// 2. 一次满足情况产生，在这个情况下定位min start end
		for match == len(need) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}

			// 3. 触发处理left++事件
			c = t[left]
			left++
			// 4. left++会改变win和 match
			if need[c] == win[c] {
				match--
			}
			win[c]--
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}
