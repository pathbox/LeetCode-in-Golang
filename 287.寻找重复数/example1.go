package LeetCode287

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	slow, fast = nums[slow], nums[nums[fast]]
	// 进入环内
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	// 找到环的入口
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return fast
}
