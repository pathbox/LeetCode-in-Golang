package LeetCode026

// 快慢索引
func removeDuplicates(nums []int) int {
	low := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[low] != nums[fast] { // 如果不相等，其实两个索引都会往后移动，如果相等，则慢索引保持不动，快索引继续往后移动
			nums[low+1] = nums[fast]
			low++
		}
	}
	return low + 1
}

// 如果相等，此时low不动，fast继续移动，此时的nums[low]==nums[low+1], nums[low+1]这个值是要被移除的，它就等着fast往后移动找到不相等的一个值，然后nums[low+1] = nums[fast]，low+1上的值就被后面不相等的值覆盖了
