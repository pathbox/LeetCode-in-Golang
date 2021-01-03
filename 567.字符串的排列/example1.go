package LeetCode567

func checkInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if cnt1 == cnt2 { // 两个窗口相等
			return true
		}
		// cnt2窗口右移
		cnt2[s2[i]-'a']--
		cnt2[s2[i+len(s1)]-'a']++ // 窗口大小就是len(s1)
	}
	return cnt1 == cnt2 // 最后一次的窗口还需要判断一下
}
