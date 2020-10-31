package LeetCode986

func intervalIntersection(A [][]int, B [][]int) [][]int {
	i, j := 0, 0
	res := make([][]int, 0)
	for i < len(A) && j < len(B) { // 循环的条件
		a1, a2 := A[i][0], A[i][1]
		b1, b2 := B[j][0], B[j][1]
		if a2 >= b1 && b2 >= a1 { //有交集区间的条件
			tmp := []int{max(a1, b1), min(a2, b2)}
			res = append(res, tmp)
		}
		if b2 < a2 { // y谁更小，索引增加取下一个，更大的保持不变
			j++
		} else {
			i++
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
