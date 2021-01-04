package LeetCode850

import "sort"

func rectangleArea(rectangles [][]int) int {
	const mod = 1e9 + 7

	xSplit, ySplit := make([]int, 0, len(rectangles)), make([]int, 0, len(rectangles))
	for _, v := range rectangles {
		xSplit = appendNoDouble(xSplit, v[0])
		ySplit = appendNoDouble(ySplit, v[1])
		xSplit = appendNoDouble(xSplit, v[2])
		ySplit = appendNoDouble(ySplit, v[3])
	}
	sort.Ints(xSplit)
	sort.Ints(ySplit)
	area := 0
	for i := 0; i+1 < len(xSplit); i++ {
		for j := 0; j+1 < len(ySplit); j++ {
			area += getArea(rectangles, xSplit[i], ySplit[j], xSplit[i+1], ySplit[j+1])
			area %= mod
		}
	}
	return area

}

func getArea(rectangles [][]int, x1 int, y1 int, x2 int, y2 int) int {
	for _, v := range rectangles {
		if v[0] <= x1 && v[1] <= y1 && v[2] >= x2 && v[3] >= y2 {
			return (x2 - x1) * (y2 - y1)
		}
	}
	return 0
}

func appendNoDouble(split []int, i int) []int {
	for _, v := range split {
		if v == i {
			return split
		}
	}
	return append(split, i)
}
