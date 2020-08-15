package LeetCode240

// Time: O(m+n), Space: O(1), m和n分别是二维数组的行数和列数
func searchMatrix(matrix [][]int, target int) bool {
	// 检查所有边界情况
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		return false // 如果出现直接返回[-1，-1]
	}
	// 然后计算出行数m和列数n
	m, n := len(matrix), len(matrix[0])
	// 初始化游标i为0，j为n-1，表示指向二维数组右上角的数字
	i, j := 0, n-1 // 从右上角开始
	for i < m && j >= 0 {
		// 如果目标值小于右上角的数字
		if target < matrix[i][j] {
			j-- // 列下标-1，排除这一列数字，向左移动一位
		} else if target > matrix[i][j] {
			i++ // 行坐标+1，排序这一行数字，向下移动一位
		} else {
			return true // 如果相等
		}
	}
	// 循环结束后如果还没有找到目标值，返回[-1,-1]
	return false
}
