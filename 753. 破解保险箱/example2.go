package LeetCode753

import (
	"math"
	"strings"
)

func crackSafe(n int, k int) string {
	total := int(math.Pow(float64(k), float64(n)))
	high := total / k
	seen := make([]bool, total)
	buf := strings.Builder{}
	var dfs func(cur int)
	dfs = func(cur int) {
		if seen[cur] {
			return
		}
		seen[cur] = true
		last := cur % k          // cur 的最后一位，要加入结果
		cur = cur % high * k     // cur 删除最高位，最后一位补 0
		for i := 0; i < k; i++ { // 尝试在最低位追加 i
			next := cur + i
			dfs(next)
		}
		// 后序遍历，追加 cur 的最后一位
		buf.WriteByte(byte(last) + '0')
	}
	dfs(0)

	// 一开始的n-1 位 0，因后序遍历的原因，这些 0 也在随后追加
	for i := 1; i < n; i++ {
		buf.WriteByte('0')
	}
	// 实际应该返回逆序，但是对这个问题，正序、逆序都行
	return buf.String()
}
