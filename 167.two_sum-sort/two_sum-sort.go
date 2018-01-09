package Problem167

// package main

// two pointer
func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		s := nums[l] + nums[r]
		if s == target {
			return []int{l + 1, r + 1}
		} else if s < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}

// func main() {
// 	nums := []int{2, 7, 11, 15}
// 	target := 9

// 	result := twoSum(nums, target)
// 	fmt.Println("result:", result)
// }
