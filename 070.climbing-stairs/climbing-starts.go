package LeetCode070

func climbStairs(n int) int {

	if n < 2 {
		return 1
	}

	rec := make([]int, n+1)
	rec[0], rec[1] = 1, 1

	for i := 2; i <= n; i++ {
		rec[i] = rec[i-1] + rec[i-2] // 非递归方式， i 从2开始将每个对应的可能情况都放到对应的rec[i]
		// rec[n] = rec[n-1]+rec[n-2]
	}

	return rec[n]
}

// 递归方式
func climbStairs_recursion(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return 1
	}

	r := climbStairs(n-1) + climbStairs(n-2)
	return r
}
