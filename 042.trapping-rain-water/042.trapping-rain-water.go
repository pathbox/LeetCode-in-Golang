package LeetCode0042

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0 // 至少要三个点菜能储水
	}

	// left[i] 是 height[:i+1] 中的最大值
	// right[i] 是 height[i:] 中的最大值
	left, right := make([]int, length), make([]int, length)

	left[0], right[length-1] = height[0], height[length-1]

	for i := 1; i < length; i++ {
		left[i] = max(left[i-1], height[i])
		right[length-1-i] = max(right[length-i], height[length-1-i])
	}

	water := 0
	for i := 0; i < length; i++ {
		// 存水量取决于 左右最大值 中的较小值
		water += min(left[i], right[i]) - height[i]
	}

	return water

}
