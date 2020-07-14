package LeetCode217

import "sort"

func containsDuplicate(nums []int) bool {
	// 将数组排序
	sort.Ints(nums)
	// 将有序数组进行遍历 将循环变量赋为1是为了避免数组下标越界这个所悟
	for i := 1; i < len(nums); i++ {
		/*
		   将当前变量与前一个变量进行比较
		   如果想要提升时间
		   也可以使用^运算符 遇相同的位将置为0
		   nums[i] ^ nums[i-1] == 0
		*/
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
