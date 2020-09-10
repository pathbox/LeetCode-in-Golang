package Offer066

func constructArr(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	size := len(a)
	left, right := make([]int, size), make([]int, size)

	left[0] = 1
	for i := 1; i < size; i++ {
		left[i] = left[i-1] * a[i-1]
	}

	right[size-1] = 1
	for i := size - 2; i >= 0; i-- {
		right[i] = right[i+1] * a[i+1]
	}

	b := make([]int, size)
	for i := 0; i < size; i++ {
		b[i] = left[i] * right[i]
	}

	return b
}
