package main

import "fmt"

func merge(ary []int, l, r, mid int) {
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = ary[i]
	}

	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		if left > mid {
			ary[i] = temp[right-l]
			right++
		} else if right > r {
			ary[i] = temp[left-l]
			left++
		} else if temp[left-l] > temp[right-l] {
			ary[i] = temp[right-l]
			right++
		} else {
			ary[i] = temp[left-l]
			left++
		}
	}
}

func MergeSort(ary []int, l, r int) {
	// 第二步优化，当数据规模足够小的时候，可以使用插入排序
	if r-l <= 15 {
		// 对 l,r 的数据执行插入排序
		for i := l + 1; i <= r; i++ {
			temp := ary[i]
			j := i
			for ; j > 0 && temp < ary[j-1]; j-- {
				ary[j] = ary[j-1]
			}
			ary[j] = temp
		}
		return
	}

	mid := (r + l) / 2
	MergeSort(ary, l, mid)
	MergeSort(ary, mid+1, r)

	// 第一步优化，左右两部分已排好序，只有当左边的最大值大于右边的最小值，才需要对这两部分进行merge操作
	if ary[mid] > ary[mid+1] {
		merge(ary, l, mid, r)
	}
}

func main() {
	ary := []int{3, 1, 2, 5, 6, 43, 4, 12, 15, 11, 55, 34}
	MergeSort(ary, 0, len(ary)-1)

	fmt.Println(ary)
}
