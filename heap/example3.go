package main

import "fmt"

// 先自底向上构建最大堆，再移除堆元素实现堆排序
func HeapSort(array []int) {
	count := len(array)

	start := count/2 + 1 // 最底层的父节点
	end := count - 1

	// 1.从最底层开始，逐一对节点进行下沉
	for start >= 0 {
		sift(array, start, count)
		start-- // 表示左偏移一个节点，如果该层没有节点了，那么表示到了上一层的最右边
	}

	// 2.下沉结束了，现在要来排序了
	// 元素大于2个的最大堆才可以移除
	for end > 0 {
		// 将堆顶元素与堆尾元素互换，表示移除最大堆元素
		array[end], array[0] = array[0], array[end]
		// 对堆顶进行下沉操作
		sift(array, 0, end)
		// 一直移除堆顶元素
		end--
	}
}

// 下沉操作，需要下沉的元素时 array[start]，参数 count 只要用来判断是否到底堆底，使得下沉结束
func sift(array []int, start, count int) {
	root := start
	// 左儿子
	child := root*2 + 1
	for child < count {
		// 右儿子比左儿子大，那么要翻转的儿子改为右儿子
		if count-child > 1 && array[child] < array[child+1] {
			child++
		}
		// 父亲节点比儿子小，那么将父亲和儿子位置交换
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			// 继续往下沉
			root = child
			child = root*2 + 1
		} else {
			return
		}
	}
}
func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}

	HeapSort(list)

	// 打印排序后的值
	fmt.Println(list)
}

/*
// siftDown implements the heap property on data[lo:hi].
// first is an offset into the array where the root of the heap lies.
func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}
*/
