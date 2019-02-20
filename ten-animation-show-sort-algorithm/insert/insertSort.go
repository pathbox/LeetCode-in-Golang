package insert

func InsertSort(ary []int) []int {
	l := len(ary)
	// 从下标为1的元素开始选择合适的位置插入，因为下标为0的只有一个元素，默认是有序的
	for i := 1; i < l; i++ {
		tmp := ary[i]
		// 从已经排序的序列最右边的开始比较，找到比其小的数
		j := i
		for j > 0 && tmp < ary[j-1] { // j > 0: 遍历到第二个元素就行
			ary[j] = ary[j-1]
			j--
		}

		if j != i {
			ary[j] = tmp
		}
	}

	return ary
}
