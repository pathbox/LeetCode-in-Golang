package Offer059

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var max int

	for i, j := 0, k-1; j >= 0 && j < len(nums); j++ {
		if i == 0 || max == nums[i-1] {
			max = nums[i]
			for t := j; t > i; t-- {
				if max < nums[t] {
					max = nums[t]
				}
			}
		} else if nums[j] > max {
			max = nums[j]
		}
		result = append(result, max)
		i++
	}
	return result
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int

	for i, j := 0, k-1; j >= 0 && j < len(nums); j++ {
		max := nums[i]
		for k := j; k > i; k-- {
			if max < nums[k] {
				max = nums[k]
			}
		}
		result = append(result, max)
		i++
	}
	return result
}
