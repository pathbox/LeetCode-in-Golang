package LeetCode056

import "sort"

func merge(intervals [][]int) [][]int {
	ret := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] // 对 x坐标从小到大排序
	})
	var cs []int // 上一个合并后的临时元素
	for _, s := range intervals {
		if len(ret) == 0 || cs[1] < s[0] { // 不能合并
			ret = append(ret, s)
			cs = s
		} else if cs[1] < s[1] { // 能合并，just 扩展y坐标即可
			cs[1] = s[1]
		}
	}
	return ret
}
