package LeetCode973

import "sort"

type Node struct {
	val   []int
	count int
}

func kClosest(points [][]int, K int) [][]int {
	n := make([]Node, 0, len(points))
	for _, v := range points {
		n = append(n, Node{v, v[0]*v[0] + v[1]*v[1]})
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i].count < n[j].count
	})
	ans := make([][]int, 0, K)
	n = n[:K]
	for _, v := range n {
		ans = append(ans, v.val)
	}
	return ans
}
