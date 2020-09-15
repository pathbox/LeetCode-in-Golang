package LeetCode048

func rotate(matrix [][]int) {
	length := len(matrix)

	// 先左<->右互换
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-1-j] = matrix[i][length-1-j], matrix[i][j]
		}
	}
	// 后关于左下/右上的对角线互换
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			matrix[i][j], matrix[length-j-1][length-i-1] = matrix[length-j-1][length-i-1], matrix[i][j]
		}
	}
}
