package LeetCode452

import "sort"

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return n
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	pos := points[0][1]
	count := 1
	for i := 1; i < n; i++ {
		if points[i][0] <= pos {
			if points[i][1] < pos { // 与pos完全重合 pos是更小的那个1索引值,如果pos < points[i][1] 不用重设pos
				pos = points[i][1]
			}
		} else { // 说明没有重复区域，需要多一根箭
			count++
			pos = points[i][1]
		}
	}
	return count
}
