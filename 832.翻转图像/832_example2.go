package LeetCode832

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		l := 0
		r := len(A[i]) - 1
		for l <= r {
			if l < r {
				A[i][l], A[i][r] = A[i][r], A[i][l]
				A[i][l] = A[i][l] ^ 1 // 异或
				A[i][r] = A[i][r] ^ 1
			} else {
				A[i][l] = A[i][l] ^ 1 // 当l=r的情况
			}
			l++
			r--
		}
	}
	return A

}
