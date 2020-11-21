package LeetCode1724

import "math"

func getMaxMatrix(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	n := len(matrix)
	m := len(matrix[0])
	res := math.MinInt32
	var left, right, up, down int
	for i := 0; i < n; i++ {
		arr := make([]int, m)
		for j := i; j < n; j++ {
			combine(arr, matrix[j]) // arr 就是累加得到的一维数组
			sum := 0
			var l int
			for t, v := range arr { // 对一维数据进行求最大前缀和
				if sum < 0 { // 如果sum<0则会重置 sum 以及左边节点坐标
					sum = 0
					l = t
				}
				sum += v
				if res < sum {
					res = sum                          // res 就是得到的最大前缀和
					left, right, up, down = l, t, i, j // 上下左右四个边界
				}
			}
		}
	}
	return []int{up, left, down, right}
}

// 叠加
func combine(arr1, arr2 []int) {
	for i, v := range arr2 {
		arr1[i] += v
	}
}
