package LeetCode378

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if check(matrix, mid, k, n) {
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
	num := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			num += i + 1 // 有更小的值统计+1
			j++
		} else {
			i--
		}
	}
	return num >= k
}
