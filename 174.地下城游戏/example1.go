package LeetCode174

import "math"
// f(0,0)+dungeon[0][0]=min(f(0+1,0),f(0,0+1)) dp[i][j] 表示从坐标 (i,j)(i,j) 到终点所需的最小初始值 最小的初始值是不会为负数的
func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[n][m-1], dp[n-1][m] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			minn := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(minn-dungeon[i][j], 1) //反向 dp dp[i][j]=max{1,min{dp[i+1][j],dp[i][j+1]}−dungeon[i][j]}
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//  时间复杂度：O(N×M)，其中 N,MN,M 为给定矩阵的长宽。

// 空间复杂度：O(N×M)，其中 N,MN,M 为给定矩阵的长宽，注意这里可以利用滚动数组进行优化，优化后空间复杂度可以达到 O(N)。
