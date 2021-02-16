package LeetCode753

import (
	"math"
	"strconv"
)

func crackSafe(n int, k int) string {
	seen := map[int]bool{}
	ans := ""
	highest := int(math.Pow(10, float64(n-1)))

	var dfs func(int)
	dfs = func(node int) {
		for x := 0; x < k; x++ {
			nei := node*10 + x
			if !seen[nei] {
				seen[nei] = true
				dfs(nei % highest)
				ans += strconv.Itoa(x)
			}
		}
	}

	dfs(0)
	for i := 0; i < n; i++ {
		ans += "0"
	}
	return ans
}
