package LeetCode560

func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	var sum int
	var count int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			count += v
		}
		m[sum]++
	}
	return count
}
