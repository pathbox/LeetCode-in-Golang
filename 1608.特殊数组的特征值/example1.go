package LeetCode1608

func specialArray(nums []int) int {
	counting := make([]int, 1001)
	for i := range nums {
		counting[nums[i]]++
	}

	preSum := 0
	for i := 1000; i >= 0; i-- {
		preSum += counting[i]
		if preSum == i {
			return i
		}
	}
	return -1
}
