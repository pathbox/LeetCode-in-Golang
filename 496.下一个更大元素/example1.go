package LeetCode496

/*
时间复杂度：O(M+N)O(M+N)，其中 MM 和 NN 分别是数组 nums1 和 nums2 的长度。

空间复杂度：O(N)O(N)。我们在遍历 nums2 时，需要使用栈，以及哈希映射用来临时存储答案。

*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)  // val 是key右边比其大的值
	stack := make([]int, 0) // 单调递减栈

	for i := 0; i < len(nums2); i++ {
		for len(stack) != 0 && nums2[i] > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			m[top] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	res := make([]int, 0)
	for _, v := range nums1 {
		if v, ok := m[v]; ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}
	}
	return res
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var stack []int
	for _, v := range nums2 {
		for len(stack) != 0 && v > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
	}

	for k, v := range nums1 {
		if v, ok := m[v]; ok {
			nums1[k] = v
		} else {
			nums1[k] = -1
		}
	}
	return nums1
}
