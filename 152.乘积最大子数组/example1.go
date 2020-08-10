package LeetCode152

func maxProduct(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	imin := 1
	imax := 1
	mx := ^(int(^uint(0) >> 1))

	for i := 0; i < len(nums); i++ {
		tmp := imax
		imax = max(max(nums[i], imax*nums[i]), imin*nums[i])
		imin = min(min(nums[i], imin*nums[i]), tmp*nums[i])
		mx = max(mx, imax)
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/*
func maxProduct(nums []int) int {
    temp0, temp1, res := nums[0], nums[0], nums[0]
	for i:=0; i<len(nums)-1; i++ {
		if nums[i+1] >= 0 {
			temp0, temp1 = max(temp0 * nums[i+1], nums[i+1]), min(temp1 * nums[i+1], nums[i+1])
		} else {
			temp0, temp1 = max(temp1 * nums[i+1], nums[i+1]), min(temp0 * nums[i+1], nums[i+1])
		}
		if res < temp0 {
			res = temp0
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b{
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b{
		return b
	}
	return a
}

*/
