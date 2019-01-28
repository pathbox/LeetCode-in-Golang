package selectSort

func SelectSort(ary []int) []int {
	l := len(ary)

	for i := 0; i < l-1; i++ {
		min := i

		// 每轮需要比较的次数 N-i
		for j := i + 1; j < l; j++ {
			if ary[j] < ary[min] {
				min = j // 记录目前能找到的最小值元素的下标
			}
		}
		// 将找到的最小值和i位置所在的值进行交换
		if i != min {
			tmp := ary[i]
			ary[i] = ary[min]
			ary[min] = tmp
		}
	}
	return ary
}
