package Offer040

func getLeastNumbers(arr []int, k int) []int {
	// 返回排序后的0-k个元素
	return quickSort(arr[:])[:k]
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// 直接取第一个元素为基准值
	pivot := arr[0]

	// 初始化变量，left表示小于基准值，middle表示和基准值一样的元素，right表示大于基准值的元素
	var left, middle, right []int
	for _, v := range arr {
		// 小于基准值的放左边，等于的放中间，大于的放右边
		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			middle = append(middle, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}
	// 递归调用此方法，直至数组全部排序完毕
	return append(append(quickSort(left[:]), middle...), quickSort(right[:])...)
}
