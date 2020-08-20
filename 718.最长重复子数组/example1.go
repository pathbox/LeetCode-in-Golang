package LeetCode718

/*
dp[i][j] ：长度为 ii，末尾项为 A[i-1]A[i−1] 的子数组，长度为 jj，末尾项为 B[j-1]B[j−1] 的子数组，二者的最大公共后缀子数组长度。（即以 A[i-1]A[i−1]（B[j-1]B[j−1]）为末尾项的公共子数组）
如果 A[i-1] != B[j-1]， dp[i][j] = 0
如果 A[i-1] == B[j-1] ， dp[i][j] = dp[i-1][j-1] + 1
*/
func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	max := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max
}
