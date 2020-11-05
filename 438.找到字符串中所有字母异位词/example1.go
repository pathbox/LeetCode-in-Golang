package LeetCode0438

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	lp := len(p)
	ls := len(s)
	if lp > ls {
		return res
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range p {
		cnt1[p[i]-'a']++
		cnt2[s[i]-'a']++
	}

	for i := 0; i <= ls-lp; i++ {
		if cnt1 == cnt2 { // 两个窗口相等
			res = append(res, i)
		}
		cnt2[s[i]-'a']-- // 左边窗口移动
		// 右边窗口移动
		if i+lp < ls {
			cnt2[s[i+lp]-'a']++ // 窗口大小就是lp
		}
	}
	return res
}
