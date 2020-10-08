package LeetCode042

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	size := len(height)
	res := 0
	leftDP := make([]int, size)  // 0-i最大的值
	rightDP := make([]int, size) // i-size最大的值
	leftDP[0] = height[0]
	rightDP[size-1] = height[size-1]
	for i := 1; i < size; i++ {
		leftDP[i] = max(leftDP[i-1], height[i])
	}
	for i := size - 2; i >= 0; i-- {
		rightDP[i] = max(rightDP[i+1], height[i])
	}
	for i := 1; i < size-1; i++ {
		minHeight := min(leftDP[i], rightDP[i])
		water := minHeight - height[i]
		res += water
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
	if a > b {
		return b
	}
	return a
}
