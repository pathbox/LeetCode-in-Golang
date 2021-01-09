// https://www.bilibili.com/video/BV15D4y1X7Tt?p=9

package Offer0109

func QueryBox(arr []int, L, R, value int) int {
	// key => num, value => idx 数组(会是一个有序数组:二分查找)
	var numM map[int][]int
	for idx, num := range arr {
		if _, ok := numM[num]; ok {
			numM[num] = append(numM[num], idx)
		} else {
			numM[num] = make([]int, 0)
			numM[num] = append(numM[num], idx)
		}
	}

	if len(numM[value]) == 0 { // 如果value 不存在arr中
		return 0
	}

	idxArr := numM[value]
	// 查询<L的下标有几个
	a := countLess(idxArr, L)
	// 查询<R的下标有几个
	b := countLess(idxArr, R)
	return b - a
}

// 因为idxArr是有序的数组，使用二分法去得到结果 找最右边小于n的数
// 在有序数组arr中，用二分法数出<n的数有几个，也就是二分法，找到<n的数中最右的位置
func countLess(arr []int, n int) int {
	l := 0
	r := len(arr) - 1
	mostRight := -1

	for l <= r {
		mid := l + (r-l)>>1
		if arr[mid] < n {
			mostRight = arr[mid]
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return mostRight
}
