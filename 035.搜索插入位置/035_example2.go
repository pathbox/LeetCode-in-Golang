package LeetCode035

func searchInsert(nums []int, target int) int {
	n := len(nums)
	lo := 0
	hi := n - 1
	mid := 0
	for lo <= hi {
		mid = (lo + hi) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

// 二分查找
