package LeetCode075

// 遍历数组 不断的将0 值交换到首部， 2值交换到尾部，中间的自然会是1
func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for i := 0; i <= r; i++ {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
		if nums[i] == 2 {
			nums[r], nums[i] = nums[i], nums[r]
			i-- // 这里要注意哦 置换2后，i坐标-- 不进行置换，再继续比较
			r--
		}
	}
}
