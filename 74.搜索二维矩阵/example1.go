package LeetCode74

/*
将二维数组转为一维数组的方式，进行二分查找
时间复杂度 : 由于是标准的二分查找，时间复杂度为O(log(mn))。
空间复杂度 : O(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)    // 行数
	n := len(matrix[0]) // 列数

	l, r := 0, m*n-1
	for l <= r {
		pivotIdx := (l + r) / 2
		pivotVal := matrix[pivotIdx/n][pivotIdx%n] // row=pivotIdx/n col=pivotIdx%n
		if pivotVal == target {
			return true
		} else {
			if target < pivotVal {
				r = pivotIdx - 1
			} else {
				l = pivotIdx + 1
			}
		}
	}
	return false
}
