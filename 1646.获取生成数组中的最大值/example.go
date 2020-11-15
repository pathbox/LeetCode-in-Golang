package LeetCode1646

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	nums := make([]int, 100)
	res := 0
	nums[0] = 0
	nums[1] = 1

	for i := 0; i < n; i++ {
		nums[2*i] = nums[i]
		nums[2*i+1] = nums[i] + nums[i+1]
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
