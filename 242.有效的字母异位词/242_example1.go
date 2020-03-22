package LeetCode242

import "fmt"

// 1.用map 用空间换时间，
// 2.将两个字符串排序后比较是否相等，这样需要排序的时间消耗
// 3. 统计每个字符串中小写字母出现的次数，但这样要26的存储长度，也许还不如用map

func isAnagram(s, t string) bool {
	a := [26]int{}
	b := [26]int{}

	for _, v := range s {
		a[v-'a'] += 1
	}

	for _, v := range t {
		b[v-'a'] += 1
	}

	fmt.Printf("a=%v\nb=%v\n", a, b)

	return a == b

}
