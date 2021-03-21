package LeetCode039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	dfs(candidates, nil, target, 0, &res) // 深度优先
	return res
}

func dfs(candidates, nums []int, target, start int, res *[][]int) {
	if target == 0 { // 满足条件 结算
		tmp := make([]int, len(nums))
		copy(tmp, nums) //深度拷贝
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if target < candidates[i] { // 剪枝 说明之后的数不符合组合结果
			return
		}
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i, res) // 分支 传到下一个分支的dfs 仍然是i 而不是i+1,这样表示当前i在重复使用, 不用担心，上面的剪枝条件会把不满足条件的去除
	}
}
