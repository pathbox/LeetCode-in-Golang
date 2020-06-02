package offer003

// 用map
func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{}, 0)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = struct{}{}
		}
	}
	return 0
}

// 时间空间复杂度都是O(n)
