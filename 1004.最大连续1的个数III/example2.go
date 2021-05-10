package LeetCode1004

func longestOnes(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	left, right, maxCount, result := 0, 0, 0, 0
	count := make(map[int]int)

	for right < len(A) {
		c := A[right]
		if c == 1 {
			count[c]++
			maxCount = max(maxCount, count[c])
		}

		if right-left+1-maxCount > K {
			count[A[left]]--
			left++
		}
		result = max(result, right-left+1)
		right++
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
