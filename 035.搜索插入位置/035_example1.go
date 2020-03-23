package LeetCode035

func searchInsert(nums []int, target int) int {
	numsIndex := 0
	for len(nums) >= 1 {
		numsLen := len(nums)
		mid := numsLen / 2
		if nums[mid] == target {
			numsIndex += mid
			return numsIndex
		} else if nums[mid] > target {
			nums = nums[:mid]
		} else {
			numsIndex += len(nums[:mid+1]) // 每次mid的右部，numsIndex需要加上一半长度的索引值
			nums = nums[mid+1:]
		}
	}
	return numsIndex
}
