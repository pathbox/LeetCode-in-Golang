package Offer004

// 利用数组规律
// Go：从左下角开始找，利用递增关系，大于往右边找，小于往上找，超出返回false

func findNumberIn2DArray(matrix [][]int, target int) bool {
	//以左下角为原点
	i := len(matrix) - 1
	j := 0

	for i > -1 {
		if j < len(matrix[i]) {
			if target < matrix[i][j] {
				i-- // 小于target向上查找
			} else if target > matrix[i][j] {
				j++ // 大于target向右查找
			} else if target == matrix[i][j] {
				return true
			}
		} else {
			return false // 超出数组返回
		}
	}
	return false
}
