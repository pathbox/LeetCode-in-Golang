package LeetCode542

/*
定义dp[i][j]表示该位置距离0的最短距离，其实很容易想到dp[i][j]要么等于0，要么等于min(dp[i-1][j],dp[i+1][j],dp[i][j-1],dp[i][j+1])+1

这个问题棘手的就是，我们更新状态的时候，要么从左上到右下，要么右下到左上，或者更不常见的右上到左下以及左下到右上。无论哪种更新方式都只能依赖两个前置状态（比如从左上到右下时， dp[i][j]只能依赖dp[i-1][j]和dp[i][j-1]）。

这里做两遍dp状态的更新来解决上述问题， 第一次从左上到右下，转移方程为：
dp[i][j] = 0 或
dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1)
第二次更新从右下到左上，转移方程为
dp[i][j] = 0 或
dp[i][j] = min(dp[i][j], dp[i][j+1]+1, dp[i+1][j]+1)
O(mn) O(1)
*/
func updateMatrix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][j] = 0
				continue
			}
			matrix[i][j] = minTwo(getNumber(i, j-1, matrix), getNumber(i-1, j, matrix)) + 1
		}
	}

	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			matrix[i][j] = minTwo(minTwo(getNumber(i, j+1, matrix), getNumber(i+1, j, matrix))+1, matrix[i][j])
		}
	}
	return matrix
}

func getNumber(i, j int, matrix [][]int) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return 10000
	}

	return matrix[i][j]
}

func minTwo(a, b int) int {
	if a > b {
		return b
	}
	return a
}
