package LeetCode015

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums) // 首先排个序
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不同
		if first > 0 && nums[first] == nums[first-1] {
			continue // first会++
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// first 是当前指针
		// 枚举
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证b的指针在c的指针左侧,说明nums[second]+nums[third] 和太大了
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着b后续的增加
			// 就不会有满足a+b+c=0 并且b<c的c了， 可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
		// nums[second]+nums[third]+nums[first] < 0 的情况可以不用写，因为外层循环first++就是这个操作
	}
	return ans
}
