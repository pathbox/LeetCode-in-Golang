package LeetCode216

import "math"

var cnt int

func combinationSum3(k int, n int) [][]int {
	situ := [9]int{}
	cnt = k
	return helper(0, n, situ)
}

func helper(ind, nokori int, situ [9]int) [][]int { //ind表示现在的索引(也就是现在有了几个数字) nokori表示现在的几个数字距离目标还有多远
	if ind == cnt {
		if nokori == 0 {
			return append([][]int{}, append([]int{}, situ[:cnt]...)) //第二个append相当于实现了py3的[:]这种深拷贝
		} else {
			return [][]int{}
		}
	}
	i := 1
	if ind != 0 {
		i = situ[ind-1] + 1
	}
	end := int(math.Min(10, float64(nokori+1))) //剪枝2. 注意Min的两个参数都得是float64类型的
	ans := [][]int{}                            //收集结果
	for i < end {
		situ[ind] = i
		ans = append(ans, helper(ind+1, nokori-i, situ)...)
		i++
	}
	return ans
}
