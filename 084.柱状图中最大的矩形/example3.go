package LeetCode084

func largestRectangleArea(heights []int) int {
	N := len(heights)
	if N == 0 {
		return 0
	}

	// 栈的简易实现
	st, pos := make([]int, N+2), 0
	push := func(v int) {
		st[pos] = v
		pos++
	}
	pop := func() int {
		pos--
		return st[pos]
	}
	top := func() int {
		return st[pos-1]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// 首尾各加一个哨兵
	get := func(i int) int {
		if i == 0 || i == N+1 {
			return 0
		}
		return heights[i-1]
	}

	// 这里才开始
	res := 0
	for i := 0; i < N+2; i++ {
		for pos > 0 && get(top()) > get(i) {
			res = max(get(pop())*(i-top()-1), res)
		}
		push(i)
	}
	return res
}
