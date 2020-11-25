package LeetCode0108

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	m := len(matrix)
	n := len(matrix[0])

	p := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				tmp := []int{i, j}
				p = append(p, tmp)
			}
		}
	}

	for _, v := range p {
		x := v[0]
		y := v[1]
		for j := 0; j < n; j++ {
			matrix[x][j] = 0
		}
		for i := 0; i < m; i++ {
			matrix[i][y] = 0
		}
	}
}
