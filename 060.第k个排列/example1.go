package LeetCode060

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/permutation-sequence/solution/golang-100-faster-by-a-bai-152/
func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	tokens := make([]int, n)
	for i := 1; i < n+1; i++ {
		factorial[i] = factorial[i-1] * i // 提前计算n!存到factorial中
		tokens[i-1] = i
	}

	k--

	var b strings.Builder
	for n > 0 {
		n--
		a := k / factorial[n]
		k = k % factorial[n]
		fmt.Fprintf(&b, "%d", tokens[a])
		tokens = append(tokens[:a], tokens[a+1:]...) // 每次a这个索引的值删除
	}
	return b.String()
}
