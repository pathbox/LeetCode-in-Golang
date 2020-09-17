package LeetCode056

import "sort"

// 上一个坐标的y坐标大于当前的x坐标并且y坐标小于当前y，表示可以合并 c[0]:x, c[1]:y
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
		} else if cs[1] >= s[0] && cs[1] < s[1] { // 能合并，just 扩展y坐标即可, cs[1] >= s[0]可以不写，因为上一个条件判断了cs[1] < s[0]操作
			cs[1] = s[1] // 能合并 改变c的y坐标
		}
	}
	return ret
}
