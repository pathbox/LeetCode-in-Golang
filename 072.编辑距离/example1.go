package LeetCode072

/*
如上所述，当我们获得 D[i][j-1]，D[i-1][j] 和 D[i-1][j-1] 的值之后就可以计算出 D[i][j]。

D[i][j-1] 为 A 的前 i 个字符和 B 的前 j - 1 个字符编辑距离的子问题。即对于 B 的第 j 个字符，我们在 A 的末尾添加了一个相同的字符，那么 D[i][j] 最小可以为 D[i][j-1] + 1；

D[i-1][j] 为 A 的前 i - 1 个字符和 B 的前 j 个字符编辑距离的子问题。即对于 A 的第 i 个字符，我们在 B 的末尾添加了一个相同的字符，那么 D[i][j] 最小可以为 D[i-1][j] + 1；

D[i-1][j-1] 为 A 前 i - 1 个字符和 B 的前 j - 1 个字符编辑距离的子问题。即对于 B 的第 j 个字符，我们修改 A 的第 i 个字符使它们相同，那么 D[i][j] 最小可以为 D[i-1][j-1] + 1。特别地，如果 A 的第 i 个字符和 B 的第 j 个字符原本就相同，那么我们实际上不需要进行修改操作。在这种情况下，D[i][j] 最小可以为 D[i-1][j-1]。

那么我们可以写出如下的状态转移方程：

若 A 和 B 的最后一个字母相同：

\begin{aligned} D[i][j] &= \min(D[i][j - 1] + 1, D[i - 1][j]+1, D[i - 1][j - 1])\\ &= 1 + \min(D[i][j - 1], D[i - 1][j], D[i - 1][j - 1] - 1) \end{aligned}
D[i][j]
​

=min(D[i][j−1]+1,D[i−1][j]+1,D[i−1][j−1])
=1+min(D[i][j−1],D[i−1][j],D[i−1][j−1]−1)
​


若 A 和 B 的最后一个字母不同：

D[i][j] = 1 + \min(D[i][j - 1], D[i - 1][j], D[i - 1][j - 1])
D[i][j]=1+min(D[i][j−1],D[i−1][j],D[i−1][j−1])
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 空字符串编辑为空字符串的编辑代价为0
	dp[0][0] = 0
	// dp[i][0]表示编辑为空串的代价，即为将所有字符删除的代价
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	// dp[0][i]表示将空串编辑为str2[:i]的代价，即插入字符的代价
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	// 下面是动态规划的主方法 所以从1开始
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				// dp[i][j] = dp[i-1][j-1] + 1 // 替换代价
			}
			// dp[i][j] = min(dp[i][j], dp[i][j-1]+1) // 插入代价
			// dp[i][j] = min(dp[i][j], dp[i-1][j]+1) // 删除代价 i-1已经和j对应上了，所以i和j对应的时候，就是要一次删除操作删掉i
		}
	}
	return dp[m][n]
}

func min(args ...int) int {
	m := args[0]
	for _, v := range args {
		if m > v {
			m = v
		}
	}
	return m
}

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }
