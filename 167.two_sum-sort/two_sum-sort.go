// package Problem167
package main

import "fmt"

// two pointer
func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1 // 获得头尾两端的数组索引

	for l < r { // 不断进行循环
		s := nums[l] + nums[r]
		if s == target { // 相等，说明找到了
			return []int{l + 1, r + 1}
		} else if s < target { // 如果和小于target，说明和不够，应该头部往后走，以达到增大和的目的
			l++
		} else { // 如果和大于target，说明和太大，应该尾部往前走，以达到减小和的目的
			r--
		}
	}
	return []int{}
}

func main() {
	nums := []int{1, 2, 3, 4, 4, 9, 56, 90}
	target := 8

	result := binarySearchTwoSum(nums, target)
	fmt.Println("result:", result)
}

// binary search
func binarySearchTwoSum(nums []int, target int) []int {
	for i, _ := range nums {
		l, r := i, len(nums)-1
		tmp := target - nums[i]

		for l <= r {
			mid := l + (r-l)>>1
			if nums[mid] == tmp {
				if mid == i {
					if nums[i]+nums[mid+1] == target {
						return []int{i + 1, mid + 1 + 1}
					}
				}
				return []int{i + 1, mid + 1}
			} else if nums[mid] < tmp {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return []int{}
}
