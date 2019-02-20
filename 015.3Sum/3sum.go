// package LeetCode15
package main

import (
	"fmt"
	"sort"
)

// func threeSum(nums []int) [][]int {

// 	numLen := len(nums)
// 	var result [][]int

// 	if numLen < 3 {
// 		return result
// 	} else if numLen == 3 {
// 		if nums[0]+nums[1]+nums[2] == 0 {
// 			result = append(result, []int{nums[0], nums[1], nums[2]})
// 			return result
// 		}
// 	}

// 	sort.Ints(nums)
// 	fmt.Print(nums)

// 	for i := 0; i < numLen-2; {
// 		low, high := i+1, numLen-1

// 		for low < high {
// 			tmp_sum := nums[i] + nums[low] + nums[high]
// 			fmt.Println(tmp_sum)
// 			if tmp_sum == 0 {
// 				result = append(result, []int{nums[i], nums[low], nums[high]})
// 			} else if tmp_sum > 0 {
// 				high--
// 			} else {
// 				low++
// 			}
// 		}
// 		i++
// 	}

// 	return result
// }

func main() {
	s := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(s)

	fmt.Println("resutl: ", res)
}

func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		// 避免添加重复的结果
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right { // 在这个条件下不断寻找 下一组的 left，right
			s := nums[i] + nums[left] + nums[right]
			switch {
			case s < 0:
				// 和小于0，left需要变大
				left++
			case s > 0:
				// 和大于0，right需要变小
				right--
			default:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 避免重复添加left，right，得到满足条件的下一组left right，再进行比较
				left, right = next(nums, left, right)

			}
		}
	}

	return res
}

func next(nums []int, left, right int) (int, int) {
	for left < right { // 在这个条件下不断寻找 下一组的 left，right
		switch {
		case nums[left] == nums[left+1]: // left 有重复的 往→走寻找不重复的left
			left++
		case nums[right] == nums[right-1]: // right有重复的往←走寻找不重复的right
			right--
		default: // 返回不重复的 下一组 left，right
			left++
			right--
			return left, right
		}
	}
	return left, right
}
