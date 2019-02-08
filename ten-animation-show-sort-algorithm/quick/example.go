package main

import "fmt"

func main() {
	arr := []int{5, 2, 3, 8, 4, 6, 6, 11, 55, 33, 26, 109}
	fmt.Println("Before sort: ", arr)
	l, r := 0, len(arr)
	result := QuickSort(l, r-1, arr)
	fmt.Println("After sort: ", result)
}
func QuickSort(left, right int, ary []int) []int {
	var i, j, temp int
	if left > right { // 说明没有要排的数，返回
		return ary
	}

	i = left
	j = right
	temp = ary[left] // 以left为基准值

	for i != j {
		for ary[j] >= temp && i < j {
			j-- // 右边的索引开始不断移动
		}

		for ary[i] <= temp && i < j {
			i++ // 左边的索引不断移动
		}
		// 此时的 i j 对应的值需要交换， ary[i]大于等于基准值，ary[j]小于等于基准值
		if i < j {
			ary[i], ary[j] = ary[j], ary[i]
		}
	}
	//i==j 相遇了 最终将基准数归位 相遇的位置和基准值进行交换，这个位置就是这个基准值的顺序位置
	ary[left], ary[i] = ary[i], ary[left]

	// 上面的操作将此时i位置的值排好序了，下面进行递归，将i的左边和右边，分别进行相同的操作排序

	QuickSort(left, i-1, ary)
	QuickSort(i+1, right, ary)

	return ary
}
