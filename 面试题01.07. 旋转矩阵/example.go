package LeetCode0107

/*
时间复杂度：O(N^2)，其中 NN 是 \textit{matrix}matrix 的边长。对于每一次翻转操作，我们都需要枚举矩阵中一半的元素。

空间复杂度：O(1)O(1)，为原地翻转得到的原地旋转。

*/
// func rotate(matrix [][]int) {
// 	row := len(matrix)

// 	for i := 0; i < row; i++ {
// 		for j := 0; j < row; j++ {
// 			matrix[i][j], matrix[j][row-i-1] = matrix[j][row-i-1], matrix[i][j]
// 		}
// 	}
// }

func rotate(matrix [][]int) {
	n := len(matrix)
	row := n / 2
	col := (n + 1) / 2
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}
}

func rotate(matrix [][]int) {
	n := len(matrix)

	// 水平翻转
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
	//主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
