package LeetCode820

import "sort"

func minimumLengthEncoding(words []string) int {
	suffix := map[string]bool{} // 用一个map存所有后缀
	sort.Sort(ByLength(words))
	res := 0
	for _, i := range words {
		if _, ok := suffix[i]; ok {
			continue
		} else {
			res += len(i) + 1
		}
		for j := 0; j < len(i)-1; j++ { // 遍历每个单词元素，存储后缀
			suffix[i[j:]] = true
		}
	}
	return res
}

type ByLength []string

func (p ByLength) Len() int {
	return len(p)
}
func (p ByLength) Less(i, j int) bool {
	return len(p[i]) > len(p[j])
}
func (p ByLength) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
