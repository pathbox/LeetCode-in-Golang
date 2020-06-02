package offer003

// 原地置换 O(n) 空间O(1)
// 从数组首端开始扫描数组，对于下标为i的数字nums[i]，如果其等于下标i，则扫描下一个数字。如果不等，
// 则把它和下标为nums[i]的数字nums[nums[i]]比较，如果两者相等，则说明nums[i]为重复数字，返回其值即可；如果不等，就把第i个数字和第nums[i]个数字交换，这样第nums[i]个数字对应的就是nums[i],接下来再重复这个交换比较的过程。

func findRepeatNumber(nums []int) int {
	for index, num := range nums {
		if index == num {
			continue
		}
		if nums[num] == num {
			return num
		}
		nums[num], nums[index] = nums[index], nums[num]
	}
	return 0
}
