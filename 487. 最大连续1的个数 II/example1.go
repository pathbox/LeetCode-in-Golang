package LeetCode487

func findMaxConsecutiveOnes(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res, count := 0, 0
	for left, right := 0, 0; right < n; right++ {
		if nums[right] == 0 {
			count++
		}
		for count > 1 {
			if nums[left] == 0 {
				count--
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
