package LeetCode542

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

	for i := len(matrix)-1; i >=0; i-- {
		j := len(matrix[0])-1; j>=0; j--{
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
