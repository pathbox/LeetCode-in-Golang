package LeetCode378

// 外层是而分 内层是和k对比数量
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if check(matrix, mid, k, n) { // 如果次区域小于mid的数值数量大于k，说明要缩小right边界
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/solution/378java-er-fen-fa-tu-jie-you-xian-dui-lie-liang-ch/
func check(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0 // 从左下角坐标开始
	num := 0       // 在矩阵中计数有多少个元素小于等于mid，从左下角开始遍历 统计小于等于mid值的数量，比较这个数量是大于k还是小于k
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			num += i + 1 // 行坐标小于i的元素肯定都比mid小，包括matrix[i][j]一共有i + 1个
			j++
		} else {
			// 不用操作num， 因为 num += i + 1 把当前列的i以上的行都涉及到了
			i--
		}
	}
	return num >= k
}
