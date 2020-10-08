package LeetCode073

// // O(MN) 空间O(1)

func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
	if n == 0 {
		return
	}
	clearFirstR, clearFirstC := false, false // 因为第一行和第一列作为标志位了，会为1.所以需要两个标记记录原始行和列是否有0要清除的情况
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r == 0 && matrix[r][c] == 0 {
				clearFirstR = true
			}
			if c == 0 && matrix[r][c] == 0 {
				clearFirstC = true
			}
			if matrix[r][c] == 0 {
				matrix[r][0], matrix[0][c] = 0, 0
			}
		}
	}
	// handle each row
	for r := 1; r < m; r++ {
		if matrix[r][0] == 0 {
			for c := 1; c < n; c++ {
				matrix[r][c] = 0
			}
		}
	}
	for c := 1; c < n; c++ {
		if matrix[0][c] == 0 {
			for r := 1; r < m; r++ {
				matrix[r][c] = 0
			}
		}
	}
	// handle first row and first column at last
	if clearFirstR {
		for c := 1; c < n; c++ {
			matrix[0][c] = 0
		}
	}
	if clearFirstC {
		for r := 1; r < m; r++ {
			matrix[r][0] = 0
		}
	}
}
