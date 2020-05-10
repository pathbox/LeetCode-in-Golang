package LeetCode078

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	var generate2 func(pos int, num int, item []int)
	generate2 = func(pos int, num int, item []int) {
		if len(item) == num { // 满足条件，返回
			tmp := make([]int, len(item))
			copy(tmp, item)
			result = append(result, tmp)
			return
		}

		for i := pos; i < len(nums); i++ {
			item = append(item, nums[i])
			generate2(i+1, num, item)
			item = item[:len(item)-1] // 回溯
		}
	}

	for i := 0; i <= len(nums); i++ {
		item := make([]int, 0, i) // 注意cap
		generate2(0, i, item)
	}
	return result
}

/*
时间复杂度：O(N×2^N)O(N×2^N)
空间复杂度：O(N×2^N)O(N×2^N)
*/
