package main

import "fmt"

// 合并 [l,r] 两部分数据，mid 左半部分的终点，mid + 1 是右半部分的起点
func merge(arr []int, l int, mid int, r int) {
	// 因为需要直接修改 arr 数据，这里首先复制 [l,r] 的数据到新的数组中，用于赋值操作
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	// 指向两部分起点
	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		// 左边的点超过中点，说明只剩右边的数据
		if left > mid {
			arr[i] = temp[right-l]
			right++
			// 右边的数据超过终点，说明只剩左边的数据
		} else if right > r {
			arr[i] = temp[left-l]
			left++
			// 左边的数据大于右边的数据，选小的数字
		} else if temp[left-l] > temp[right-l] {
			arr[i] = temp[right-l]
			right++
		} else {
			arr[i] = temp[left-l]
			left++
		}
	}
}

func MergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	// 递归向下 不断的 /2直到一个元素位置返回
	mid := (r + l) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	// 归并向上
	merge(arr, l, mid, r)
}

func main() {
	arr := []int{3, 1, 2, 5, 6, 43, 4}
	MergeSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
