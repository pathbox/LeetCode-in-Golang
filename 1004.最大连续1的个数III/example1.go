package LeetCode1004

func longestOnes(A []int, K int) int {
	n := len(A)
	res := 0
	zeros := 0

	for left, right := 0, 0; right < n; right++ {
		if A[right] == 0 {
			zeros++
		}
		for zeros > K {
			if A[left] == 0 {
				zeros--
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
