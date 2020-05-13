package LeetCode077

var res = make([][]int, 0)

func combine(n int, k int) [][]int {
	var res = make([][]int, 0)
	helperCombine(&res, nil, 1, n, k)
	return res
}

func helperCombine(res *[][]int, data []int, start, n, k int) {
	if len(data) == k {
		*res = append(*res, data)
		return
	}
	for i := start; i <= n; i++ {
		tmp := make([]int, len(data)+1)
		copy(tmp, data)
		tmp[len(data)] = i
		helperCombine(res, tmp, i+1, n, k)
	}
}
