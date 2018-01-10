// package Problem15
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	numLen := len(nums)
	var result [][]int

	if numLen < 3 {
		return result
	} else if numLen == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			result = append(result, []int{nums[0], nums[1], nums[2]})
			return result
		}
	}

	sort.Ints(nums)
	fmt.Print(nums)

	for i := 0; i < numLen-2; {
		low, high := i+1, numLen-1

		for low < high {
			tmp_sum := nums[i] + nums[low] + nums[high]
			fmt.Println(tmp_sum)
			if tmp_sum == 0 {
				result = append(result, []int{nums[i], nums[low], nums[high]})
			} else if tmp_sum > 0 {
				high--
			} else {
				low++
			}
		}
		i++
	}

	return result
}

func main() {
	s := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(s)

	fmt.Println("resutl: ", res)
}
