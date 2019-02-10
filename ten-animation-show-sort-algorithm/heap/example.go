package main

import "fmt"

func HeapAdjust(a []int) { //调整大顶堆
	e := 0
	temp := a[0]
	for i := 1; i < len(a); i = i*2 + 1 {
		if i < len(a)-1 && a[i] < a[i+1] {
			i++
		}
		if a[i] > temp {
			a[e] = a[i]
			e = i
		} else {
			break
		}
	}
	a[e] = temp
}

func HeapInit(a []int) { //构造初始堆
	for i := len(a) / 2; i >= 0; i-- { //从第一个非叶子结点从下至上，从右至左调整结构
		HeapAdjust(a[i:])
	}
}

func HeapSort(a []int) {
	HeapInit(a)
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[0] = a[0], a[i]
		HeapAdjust(a[:i])
	}
	fmt.Println(a)
}

func main() {
	a := []int{1, 10, 19, 24, 61, 51, 121, 9, 11, 34, 121, 22}
	HeapSort(a)
}
