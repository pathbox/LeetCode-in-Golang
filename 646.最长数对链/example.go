package LeetCode646

import "sort"

type pair struct {
	data [][]int
}

func (p *pair) Len() int {
	return len(p.data)
}

func (p *pair) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

func (p *pair) Less(i, j int) bool {
	v1 := p.data[i]
	v2 := p.data[j]
	if v1[1] == v2[1] { // 如果1索引相等就比较0 要不直接比较1索引的值
		return v1[0] < v2[0]
	}
	return v1[1] < v2[1]
}

func findLongestChain(pairs [][]int) int {
	p := &pair{
		data: pairs,
	}
	sort.Sort(p)
	n := len(pairs)
	if n < 2 {
		return n
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
