package main

import "fmt"

func getMax(windowSlice []int) int {
	max := windowSlice[0]
	for i := 1; i < len(windowSlice); i++ {
		if windowSlice[i] >= max {
			max = windowSlice[i]
		}
	}
	return max

}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	//如果切片长度为0的话，那就直接返回空切片
	if len(nums) == 0 {
		return result
	}
	//如果切片长度与k一样
	if len(nums) == k {
		result = append(result, getMax(nums))
		return result
	}

	var windowSlice []int
	index := 0
	for i := k; i <= len(nums); i++ {
		windowSlice = nums[index:i]
		result = append(result, getMax(windowSlice))
		index++
	}
	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
	// nums := []int{}
	// fmt.Println(maxSlidingWindow(nums, 0))
	// nums := []int{1, -1}
	// fmt.Println(maxSlidingWindow(nums, 1))
}
