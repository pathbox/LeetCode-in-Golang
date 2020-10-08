package LeetCode073

// 遍历两边矩阵，第一遍得到row col标记位，第二遍根据标记位置换0
func setZeroes(matrix [][]int) {
	row, col := make(map[int]int), make(map[int]int)
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
