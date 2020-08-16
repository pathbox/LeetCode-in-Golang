package LeetCode268

func missingNumber(nums []int) int {
	result := 0
	for i, k := range nums {
		result ^= k ^ i
	}
	return result ^ len(nums)
}
