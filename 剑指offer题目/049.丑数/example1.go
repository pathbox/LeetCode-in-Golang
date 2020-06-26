package Offer049

// https://leetcode-cn.com/problems/chou-shu-lcof/solution/chou-shu-ii-qing-xi-de-tui-dao-si-lu-by-mrsate/
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	result := []int{1}
	res := 0
	i, j, k := 0, 0, 0
	for n > 1 {
		x := result[i] * 2
		y := result[j] * 3
		z := result[k] * 5

		res = min(min(x, y), z)

		if res == x {
			i++
		}
		if res == y {
			j++
		}
		if res == z {
			k++
		}
		result = append(result, res)
		n--
	}
	return result[len(result)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
