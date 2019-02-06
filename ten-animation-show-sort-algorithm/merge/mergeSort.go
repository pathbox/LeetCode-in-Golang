package merge

func MergeSort(ary []int) []int {
	if len(ary) < 2 {
		return ary
	}

	i := len(ary) / 2
	left := MergeSort(ary[0:i])
	right := MergeSort(ary[i:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}
