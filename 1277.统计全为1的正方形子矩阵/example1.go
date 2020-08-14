package LeetCode1227

// https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/solution/tong-ji-quan-wei-1-de-zheng-fang-xing-zi-ju-zhen-f/
func countSquares(matrix [][]int) int {
	max := 0
	// 创建二维数组作为缓存
	dp := make([][]int, len(matrix))
	// 对二维数组进行赋值
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = matrix[i][j] - 0
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			max = max + dp[i][j]
		}
	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
