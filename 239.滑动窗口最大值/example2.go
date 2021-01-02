package LeetCode239

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 || len(nums) < k {
		return []int{}
	}

	window := make([]int, 0)
	result := make([]int, 0)

	for i, v := range nums {
		if i >= k && window[0] <= i-k { // // if the left-most index is out of window, remove it 移动窗口相当于left++  left会在 i-k 到i之间
			window = window[1:]
		}

		// 单调队列处理
		for len(window) > 0 && nums[window[len(window)-1]] < v {
			window = window[0 : len(window)-1]
		}

		window = append(window, i)

		if i >= i-k {
			result = append(result, nums[window[0]]) // nums[window[0]] 是每次窗口的最大值
		}
		return result
	}

}
