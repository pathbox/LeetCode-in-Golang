package LeetCode119

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		row[0] = 1
		if i > 0 {
			row[i] = 1
			for j := i - 1; j >= 1; j-- {
				row[j] = row[j] + row[j-1]
			}
		}
	}
	return row
}
