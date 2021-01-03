package LeetCode1040

import "sort"

// https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/solution/zui-da-zui-xiao-qing-kuang-fen-kai-by-yang-wang-da/

func numMovesStonesII(stones []int) []int {
	sort.Sort(sort.IntSlice(stones)) //  升序
	n := len(stones)
	mx := stones[n-1] - stones[0] + 1 - n                       // 空格数
	mx -= min(stones[n-1]-stones[n-2]-1, stones[1]-stones[0]-1) // 最大步数=空格数-左右空挡的最小值
	var (
		mi = mx
		i  = 0
		j  = 0
	)

	for i = 0; i < n; i++ {
		for j+1 < n && stones[j+1]-stones[i]+1 <= n {
			j++
		}
		cost := n - (j - i + 1)
		if j-i+1 == n-1 && stones[j]-stones[i]+1 == n-1 {
			cost = 2
		}
		mi = min(cost, mi)
	}

	return []int{mi, mx}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
