package LeetCode040

import "sort"

// candidates 中的每个数字在每个组合中只能使用一次
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) //快排，懒得写
	res := [][]int{}
	dfs(candidates, nil, target, 0, &res) //深度优先
	return res
}

func dfs(candidates, nums []int, target, start int, res *[][]int) {
	if target == 0 { // 结算
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if i != start && candidates[i] == candidates[i-1] { // 同层节点 数值相同，剪枝 为什么要i != start呢？i = start i-1 是上一次的索引，这两者值是可以相等的
			continue
		}
		if target < candidates[i] { // 剪枝
			return
		}
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i+1, res) //*分支 i+1避免重复
	}
}

// [2,5,2,1,2], target = 5
