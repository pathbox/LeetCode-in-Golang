package radix

import (
	"math"
	"strconv"
)

func RadixSort(ary []int) []int {
	l := len(ary)
	if l <= 0 {
		return ary
	}
	max := ary[0]
	for i := 0; i < l; i++ {
		if max < ary[i] {
			max = ary[i]
		}
	}

	for j := 0; j < len(strconv.Itoa(max)); j++ {
		temp := make([][]int, 10)
		for k := 0; k < l; k++ {
			n := ary[k] / int(math.Pow(10, float64(j))) % 10
			temp[n] = append(temp[n], ary[k])
		}
		m := 0
		for p := 0; p < len(temp); p++ {
			for q := 0; q < len(temp[p]); q++ {
				ary[m] = temp[p][q]
				m++
			}
		}
	}

	return ary

}
