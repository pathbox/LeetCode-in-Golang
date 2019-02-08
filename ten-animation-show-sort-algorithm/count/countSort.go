package count

func CountSort(ary []int) []int {
	l := len(ary)

	if l == 0 {
		return ary
	}
	var max int = 0
	for i := 0; i < l; i++ {
		if ary[i] > max {
			max = ary[i] // 找到最大值
		}
	}

	var aryTemp []int = make([]int, max+1) // 创建新的数组空间,实际会浪费很多空间

	for i := 0; i < l; i++ {
		aryTemp[ary[i]]++
	}
	var j int = 0
	for i := 0; i < max+1; i++ { // i索引就是排序的数值
		if aryTemp[i] > 0 {
			ary[j] = i
			aryTemp[i]--
			j++
		}
	}

	return ary
}
