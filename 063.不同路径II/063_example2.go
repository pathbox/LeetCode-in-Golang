package LeetCode063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	for i := 0; i < m; i++ { //在初始化时，会将0行0列值为0的初始化为1，这会和障碍物的1混淆。故这里将原数组翻转一下，更清晰
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = 1
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}

	if obstacleGrid[0][0] == 0 {
		return 0
	}
	for i := 1; i < m; i++ {
		obstacleGrid[i][0] &= obstacleGrid[i-1][0]
	}
	for j := 1; j < n; j++ {
		obstacleGrid[0][j] &= obstacleGrid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}

	return obstacleGrid[m-1][n-1]
}
