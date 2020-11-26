package LeetCode1329

import "sort"

/*
其实就是把从矩阵的左侧和上侧开始往右下方走
把值拉出来排完序再放回去
*/
func diagonalSort(mat [][]int) [][]int {
	for i := 0; i < len(mat); i++ {
		t := i
		var arr []int
		for j := 0; t < len(mat) && j < len(mat[0]); j++ {
			arr = append(arr, mat[t][j])
			t++
		}
		sort.Ints(arr)
		t = i
		for j := 0; t < len(mat) && j < len(mat[0]); j++ {
			mat[t][j] = arr[j]
			t++
		}
	}
	for i := 1; i < len(mat[0]); i++ {
		t := i
		var arr []int
		for j := 0; t < len(mat[0]) && j < len(mat); j++ {
			arr = append(arr, mat[j][t])
			t++
		}
		sort.Ints(arr)
		t = i
		for j := 0; t < len(mat[0]) && j < len(mat); j++ {
			mat[j][t] = arr[j]
			t++
		}
	}
	return mat
}

// 冒泡法 O(n^3)
func diagonalSort(mat [][]int) [][]int {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			for k := 0; i+k < len(mat) && j+k < len(mat[i]); k++ {
				if mat[i][j] > mat[i+k][j+k] {
					mat[i][j], mat[i+k][j+k] = mat[i+k][j+k], mat[i][j]
				}
			}
		}
	}
}
