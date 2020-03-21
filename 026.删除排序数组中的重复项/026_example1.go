package LeetCode026

func removeDuplicates(nums []int) int {
	low := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[low] != nums[fast] {
			nums[low+1] = nums[fast]
			low++
		}
	}
	return low + 1
}
