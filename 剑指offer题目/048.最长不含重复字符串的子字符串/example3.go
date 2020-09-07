package Offer048

func lengthOfLongestSubstring(s string) int {
	res := 0
	left, right := 0, 0
	resMap := make(map[byte]int, 0)
	for right < len(s) {
		_, ok := resMap[s[right]]
		for ok { // left不断的向右缩进直到没有重复元素
			resMap[s[left]]--
			if resMap[s[left]] == 0 {
				delete(resMap, s[left])
			}
			left++
			_, ok = resMap[s[right]]
		}
		resMap[s[right]]++
		if right-left+1 > res { // 统计比较一次当前窗口子串长度
			res = right - left + 1
		}
		right++
	}
	return res
}
