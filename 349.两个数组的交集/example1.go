package LeetCode349

func intersection(nums1 []int, nums2 []int) []int {
	hash := make(map[int]bool)
	res := make([]int, 0)

	for _, v := range nums1 {
		hash[v] = true
	}

	for _, v := range nums2 {
		if hash[v] {
			res = append(res, v)
			hash[v] = false // 比较过了
		}
	}
	return res
}
