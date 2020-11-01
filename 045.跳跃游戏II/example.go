package LeetCode045

func jump(nums []int) int {
	l := len(nums)
	end, farthest := 0, 0
	jumps := 0
	for i := 0; i < l-1; i++ {
		farthest = max(nums[i]+i, farthest)
		if i == end {
			jumps++
			end = farthest
		}
	}
	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
