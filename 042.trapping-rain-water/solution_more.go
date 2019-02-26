package LeetCode042

func trap2(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}

	result := 0
	maxIdx := 0
	max := 0
	size := len(height)
	leftMax := make([]int, size)
	rightMax := make([]int, size)

	for i, v := range height {
		if v > max {
			max = v
			maxIdx = i
		}
		leftMax[i] = maxIdx // 0..i为止maxIdx
	}

	max = 0
	maxIdx = len(height) - 1
	for i := len(height) - 1; i >= 0; i-- { // 从右往左循环 i..length-1为止maxIdx
		if height[i] > max {
			max = height[i]
			maxIdx = i
		}
		rightMax[i] = maxIdx
	}

	for i, v := range height {
		if height[leftMax[i]] > v && height[rightMax[i]] > v {
			ht := min(height[leftMax[i]], height[rightMax[i]]) - v
			result += ht
		}
	}

	return result
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func trap_two_pointer(height []int) int {
	left, supermax, water := 0, 0, 0
	right := len(height) - 1
	for left < right {
		if height[left] < height[right] {
			supermax = max(height[left], supermax)
			water += supermax - height[left]
			left++
		} else {
			supermax = max(height[right], supermax)
			water += supermax - height[right]
			right--
		}
	}
	return water
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
 当从左端往右端i递增的时候，先设初始下标0为left，那么当出现height[i]大于height[left]的时候，我们就会发现，0到i之间的的积水是已经定下来了的，和后面的情况无关了，那么我们可以直接把left和i之间的积水算出来，加入积水总和sum，然后再把left更新为当前的i，重复前面的过程。

这样，我们就会最后得到一个left，这个left的右边（也就是i大于left时），不存在height[i]大于height[left]的情况，并且我们把前面所有可以得到的积水都已经加进了sum。

然后我们再从右往左，先把i和right设为len(height)-1，这是height数组的右端点，然后我们让i--，与前面同理，当height[i]>height[right]的时候，我们把之间的积水加进sum，然后把right更新为当前的i，重复这一过程，直到i碰到left或者再也没有比height[right]大的height[i]，停止。注意碰到left要停止，不能算重复了。

最后还有一个问题，left和right的最终位置有可能并不相遇，这说明在left和right之间，形成了一个洼地，也就是说对于[left,right]这个区间，两端对应的height[left]和height[right]大于区间里的所有i对应的height[i]，那么我们把这个可能的洼地里面的积水也算出来，加入sum。
*/

func trap3(height []int) int {
	lenn := len(height)
	if lenn == 0 || lenn == 1 || lenn == 2 {
		return 0
	}
	var sum int
	left := 0
	right := lenn - 1
	for i := 0; i < lenn; i++ {
		if height[left] < height[i] {
			sum += s(left, i, height)
			left = i
		}
	}
	for i := lenn - 1; i >= left; i-- {
		if height[i] > height[right] {
			sum += s(i, right, height)
			right = i
		}
	}
	sum += s(left, right, height)
	return sum
}

func s(left int, right int, height []int) int {
	sum := 0
	low_edge := height[left]
	if height[left] > height[right] {
		low_edge = height[right]
	}
	for j := 1; j < right-left; j++ {
		if low_edge > height[left+j] {
			sum += low_edge - height[left+j]
		}
	}
	return sum
}
