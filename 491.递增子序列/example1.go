package LeetCode491

// 改写别人的dfs算法
func findSubsequences(nums []int) [][]int {
	var res [][]int
	var path []int
	// cur是当前下标，length是路径长度，pre是前一个数
	dfs := func(cur, length, pre int) {}
	dfs = func(cur, length, pre int) {
		// 输出
		if length >= 2 {
			// 这里不可以直接赋值 res = append(res, path)
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		// 缓存
		cache := make(map[int]int)
		// 考虑其他数字
		for i := cur; i < len(nums); i++ {
			_, ok := cache[nums[i]]
			// 不需要
			if nums[i] < pre || ok {
				continue
			}
			path = append(path, nums[i])
			cache[nums[i]] = 1
			dfs(i+1, length+1, nums[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0, -101)
	return res
}
