package LeetCode435

import "sort"

// 首先根据二维数组的第一个数排序。

// 记录当前的末尾end， 遍历数组，如果interval[i][0] >= end 说明当前无重叠。另end = interval[i][1]
// 如果interval[i][0] < end && intervals[i][1] < end, 说明 当前end比之前的更小（更优）
// ，则剔除上一个区间（count++）， end = intervals[i][1].

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
		} else { // intervals[i][0] < end
			if intervals[i][1] < end { // 如果interval[i][0] < end && intervals[i][1] < end
				end = intervals[i][1] // 有重复区域，end取更小的那一个y端点值
			}
			count++
		}
	}
	return count
}
