package LeetCode011

func maxArea(height []int) int {
	// 从两端开始寻找，能保证了宽度是最大值
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)
		xLen := j - i
		area := h * xLen

		if area > max {
			max = area
		}

		// 朝着area具有变大的可能性方向变化
		if a < b {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
