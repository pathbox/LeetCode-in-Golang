package LeetCode073

// 遍历两边矩阵，第一遍得到row col标记位，第二遍根据标记位置换0
// O(MN) 空间O(M+N)
func setZeroes(matrix [][]int) {
	row, col := make(map[int]bool), make(map[int]bool)
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}
