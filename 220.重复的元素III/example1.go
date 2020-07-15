package LeetCode220

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1 && j-i <= k; j++ {
			if -t <= nums[j]-nums[i] && nums[j]-nums[i] <= t {
				return true
			}
		}
	}
	return false
}
