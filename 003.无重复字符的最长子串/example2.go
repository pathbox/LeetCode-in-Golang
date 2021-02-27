package LeetCode003

// 滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。一旦出现了重复字符，就需要缩小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。以此类推，每次移动需要计算当前长度，并判断是否需要更新最大长度，最终最大的值就是题目中的所求

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
			// freq[index] > 0 说明有重复值了，移动left缩小范围
		} else { // 遇到重复的字符情况，肯定是窗口的left和right+1重复(窗口的两边),所以left++
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1) // 每次计算一个窗口长度
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
