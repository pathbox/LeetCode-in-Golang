package LeetCode042

// 暴力解法 O(n^2)
func trap(height []int) int {
	l := len(height)
	res := 0
	// 起始位置是第一个，而不是0
	// 结束位置是倒数第二个，而不是倒数第一个
	// 找当前height[i]的左右两边最大
	for i := 1; i < l-1; i++ {
		leftMax := 0
		for k := 0; k <= i; k++ { // 在0-i内的左边最大
			leftMax = max(leftMax, height[k])
		}
		rightMax := 0
		for j := i; j < l; j++ { // i-l 内右边最大
			rightMax = max(rightMax, height[j])
		}
		minVal := min(leftMax, rightMax)
		water := minVal - height[i] // 减去当前的height[i]
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
