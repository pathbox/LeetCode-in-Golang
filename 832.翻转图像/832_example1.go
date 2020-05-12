package LeetCode832

func flipAndInvertImage(A [][]int) [][]int {
	for i :=0; i < len(A); i++{
		// 翻转
		for j := 0; j < len(A[i]) /2; j++{
			A[i][j],A[i][len(A[i])-1-j] = A[i][len(A[i])-1-j],A[i][j]
		}
		// 反转
		fo j :=0; j < len(A[i]);j++{
			if A[i][j] == 0 {
				A[i][j]=1
			} else {
				A[i][j] = 0
			}
		}
	}
	return A
}