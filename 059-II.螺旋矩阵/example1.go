package LeetCode059

func generateMatrix(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}

	size := n * n
	a, b := 0, n-1 //左右边界
	c, d := 0, n-1 //上下边界

	idx := 1
	for idx <= size {
		// 向右走
		for i := a; i <= b; i++ {
			result[c][i] = idx
			idx++
		}
		c++ // 上边界收缩
		// 向下走
		for i := c; i <= d; i++ {
			result[i][b] = idx
			idx++
		}
		b-- // 右边界收缩
		// 向左走
		for i := b; i >= a; i-- {
			result[d][i] = idx
			idx++
		}
		// 下边界收缩
		d--
		// 向上走
		for i := d; i >= c; i-- {
			result[i][a] = idx
			idx++
		}
		a++ // 左边界收缩
	}
	return result
}
