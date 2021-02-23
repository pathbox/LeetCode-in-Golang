package LeetCode474

func findMaxForm(strs []string, m int, n int) int {
	maap := make([][]int, 0)
	for i := 0; i <= m; i++ {
		ll := make([]int, n+1)
		maap = append(maap, ll)
	}
	one := make([]int, len(strs))
	zero := make([]int, len(strs))
	for i, _ := range strs {
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '0' {
				zero[i] += 1
			} else {
				one[i] += 1
			}
		}
	}
	for i, _ := range strs {
		for j := m; j >= zero[i]; j-- {
			for z := n; z >= one[i]; z-- {
				maap[j][z] = max(maap[j-zero[i]][z-one[i]]+1, maap[j][z])
			}
		}
	}
	return maap[m][n]

}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
