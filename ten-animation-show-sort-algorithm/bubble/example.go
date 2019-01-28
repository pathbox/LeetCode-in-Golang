package main

import "fmt"

func main() {
	arr := []int{5, 2, 3, 8, 4}
	fmt.Println("Before sort: ", arr)
	r := BubbleSort(arr)
	fmt.Println("After sort: ", r)arr := []int{5, 2, 3, 8, 4}
	fmt.Println("Before sort: ", arr)
	r := BubbleSort(arr)
	fmt.Println("After sort: ", r)
}

func BubbleSort(ary []int) []int {
	l := len(ary)

	for i := 0; i < l-1; i++ {
		// 设定一个标记，若为true，则表示此次循环没有进行交换，也就是待排序列已经有序了，排序已经完成，不需要继续下一轮
		flag := true
		for j := 0; j < l-1-i; j++ {
			if ary[j] > ary[j+1] {
				tmp := ary[j]
				ary[j] = ary[j+1]
				ary[j+1] = tmp
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return ary
}
