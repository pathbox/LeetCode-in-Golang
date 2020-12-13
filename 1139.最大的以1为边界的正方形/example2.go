package LeetCode1139

// https://leetcode-cn.com/problems/largest-1-bordered-square/solution/java-dong-tai-gui-hua-by-resolmi/

// 首先定义DP数组
// dp[i][j][0]: i,j左边连续的1的个数（包括自身）
// dp[i][j][1]: i,j上边连续的1的个数（包括自身）
func largest1BorderedSquare(grid [][]int) int {
	var dp [101][101][2]int

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			// 下面是在grid[i][j] == 1 的前提条件下
			if i == 0 {
				dp[i][j][0] = 1
			} else {
				dp[i][j][0] = dp[i-1][j][0] + 1
			}

			if j == 0 {
				dp[i][j][1] = 1
			} else {
				dp[i][j][1] = dp[i][j-1][1] + 1
			}

			// 四条“边” dp[i][j][0],dp[i][j][1],dp[i-maxLine+1][j][1],dp[i][j-maxLine+1][0]
			maxLine := min(dp[i][j][0], dp[i][j][1]) // 边长应该是较短的那一个
			// 在 maxLine > res的前提下
			for maxLine > res {
				// 另外两条边的长度也大于maxLine 说明maxLine是最合适的边长,要不然说明maxLine太长了
				if dp[i-maxLine+1][j][1] >= maxLine && dp[i][j-maxLine+1][0] >= maxLine {
					res = maxLine // 取更短的边为结果
				}
				maxLine-- // 不然maxLine是更长的边，缩小1个长度，继续寻找更短的边
			}
		}
	}
	return res * res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
