package LeetCode085

import "math"

//  https://leetcode-cn.com/problems/maximal-rectangle/solution/zui-da-ju-xing-by-leetcode/
//动态规划（太菜了，照着官方解释看懂的，估计过几天又忘了）
func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	//保存结果
	maxarea := 0
	m, n := len(matrix), len(matrix[0])
	height := make([]int, n)
	left := make([]int, n)
	right := make([]int, n)
	//初始化
	for i := 0; i < n; i++ {
		right[i] = n
	}
	for i := 0; i < m; i++ {
		cur_left, cur_right := 0, n
		//更新height
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		//更新left
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				left[j] = int(math.Max(float64(left[j]), float64(cur_left)))
			} else {
				left[j] = 0
				cur_left = j + 1
			}
		}
		//更新right
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				right[j] = int(math.Min(float64(right[j]), float64(cur_right)))
			} else {
				right[j] = n
				cur_right = j
			}
		}
		//更新最大面积
		for j := 0; j < n; j++ {
			maxarea = int(math.Max(float64(maxarea), float64(height[j]*(right[j]-left[j]))))
		}
	}
	return maxarea
}

// 时间：O(MN) O(M) M 是我们保留的额外数组的长度
