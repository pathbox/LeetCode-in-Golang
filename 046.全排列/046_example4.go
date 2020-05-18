package LeetCode047

// https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
func permute(nums []int) [][]int {
	res := [][]int{}
	used := make(map[int]bool, 0)
	path := []int{}
	dfs(nums, used, path, &res)
	return res
}

func dfs(nums []int, used map[int]bool, path []int, res *[][]int) {
	if len(nums) == len(path) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for _, v := range nums {
		if used[v] {
			continue
		}
		used[v] = true
		path = append(path, v)
		dfs(nums, used, path, res)
		used[v] = false
		path = path[:len(path)-1]
	}
}

/*
回溯法实际就是以空间换取时间的方法解决问题
首先定义used数组以及辅助path数组（used用于记录当前已用的元素 path用来辅助记录当前元素）
回溯：
结束条件：当前记录元素长度刚好等于需要全排列的个数，此时注意要重copy数组添加进结果集
回溯条件：判断当前元素是否已有，有则结束当前进入下一个循环，否则将其添加进入下一层递归
回溯还原现场：将当前元素设置为false并将path还原表示未用到，用于下一个循环
*/
