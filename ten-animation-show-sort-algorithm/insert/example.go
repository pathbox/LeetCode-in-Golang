package main

import "fmt"

func main() {
	arr := []int{5, 2, 3, 8, 4, 6, 6}
	fmt.Println("Before sort: ", arr)
	r := InsertSort(arr)
	fmt.Println("After sort: ", r)
}

func InsertSort(ary []int) []int {
	l := len(ary)
	// 从下标为1的元素开始选择合适的位置插入，因为下标为0的只有一个元素，默认是有序的
	for i := 1; i < l; i++ {
		tmp := ary[i]
		// 从已经排序的序列最右边的开始比较，找到比其小的数
		j := i
		for j > 0 && tmp < ary[j-1] { // j > 0: 每次遍历到第二个元素就行
			ary[j] = ary[j-1] // 向右平移操作，平移的过程，就是找到要去插入的位置，找到则停止循环，就行插入操作
			j--
		}

		if j != i {
			ary[j] = tmp // "插入"，其实是覆盖替换，原来的数已经平移了
		}
	}

	return ary
}
