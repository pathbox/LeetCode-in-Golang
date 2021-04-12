package main

import "fmt"

func main() {
	arr := []int{3, 4, 5, 6, 7, 8, 1, 13}
	QuickSort(0, len(arr)-1, arr)
	fmt.Println(arr)
}

func QuickSort(left, right int, array []int) {
	var i, j, temp int
	if left > right {
		return
	}

	i = left
	j = right
	temp = array[left]

	for i != j {
		for array[j] >= temp && i < j {
			j--
		}

		for array[i] <= temp && i < j {
			i++
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	//i==j 相遇了 最终将基准数归位 相遇的位置和基准值进行交换，这个位置就是这个基准值的顺序位置
	array[left], array[i] = array[i], array[left]
	QuickSort(left, i-1, array)
	QuickSort(i+1, right, array)
}
