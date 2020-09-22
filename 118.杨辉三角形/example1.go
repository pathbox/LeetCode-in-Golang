package LeetCode118

/*
构建一个二维数组切片，初始化两行
若行数小于2，则直接从切片中切取即可
从第三行开始起，每一行的第一个和最后一个都为1，中间的数计算为：row = append(row, res[i-1][j-1]+res[i-1][j])
别忘了每一行的最后一个数和将新的行添加到返回结果中
*/
func generate(numRows int) [][]int {
	var res [][]int = [][]int{
		{1},
		{1, 1},
	}
	if numRows < 2 {
		return res[0:numRows]
	}

	for i := 2; i < numRows; i++ {
		var row []int = []int{1}
		for j := 1; j < i; j++ {
			row = append(row, res[i-1][j-1]+res[i-1][j])
		}
		row = append(row, 1)
		res = append(res, row)
	}
	return res
}
