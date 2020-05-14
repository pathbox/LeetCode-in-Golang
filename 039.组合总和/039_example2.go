package LeetCode039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var result = make([][]int, 0)
	sort.Ints(candidates)
	helper(&result, []int{}, candidates, 0, target)
	return result
}

func helper(result *[][]int, tempSlice []int, candidates []int, start int, target int) {
	if target == 0 {
		*result = append(*result, tempSlice)
		return
	} else {
		for i := start; i < len(candidates); i++ {
			if target-candidates[i] >= 0 {
				newTempSlice := make([]int, len(tempSlice)+1)
				copy(newTempSlice, tempSlice)                //深度拷贝
				newTempSlice[len(tempSlice)] = candidates[i] // 当前i索引的值加入到临时slice
				// newTempSlice = append(newTempSlice,candidates[i]) // 这样不行,会导致扩容，而扩容后没有复制的索引位置的元素默认是0  [[0,2,0,2,0,3],[0,7]]
				helper(result, newTempSlice, candidates, i, target-candidates[i]) // 为什么没有回溯呢?因为没有说一定要取几个树组成一次满足的条件，也就是得到的数组是没有一个数量的界限
			} else {
				break // 剪枝
			}
		}
	}
}
