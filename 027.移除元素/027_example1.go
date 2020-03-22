package LeetCode027

func removeElement(nums []int, val int) int {
	low := 0
	fast := 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[low] = nums[fast]
			low++
		}
	}
	return low
}
