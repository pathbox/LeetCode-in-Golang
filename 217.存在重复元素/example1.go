package LeetCode217

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if m[num] {
			return true
		} else {
			m[num] = true
		}
	}

	return false
}
