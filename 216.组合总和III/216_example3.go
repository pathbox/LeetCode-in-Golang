package LeetCode216

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	if k*9 < n {
		return result
	}
	t := make([]int, 0)
	dfs4(n, k, 1, 0, t, &result)
	r := make([][]int, 0)
	for i := 0; i < len(result); i++ {
		if len(result[i]) == k {
			r = append(r, result[i])
		}
	}
	return r
}

func dfs4(n int, k int, start int, cnt int, t []int, r *[][]int) {
	if isOk(t, n) {
		tmp := make([]int, len(t))
		copy(tmp, t)
		*r = append(*r, tmp)
		return
	}
	if len(t) >= k {
		return
	}

	for i := start; i <= 9; i++ {
		t = append(t, i)
		dfs4(n, k, i+1, cnt, t, r)
		t = t[:len(t)-1]
	}

}

func isOk(t []int, n int) bool {
	var tmp int
	for i := 0; i < len(t); i++ {
		tmp = tmp + t[i]
	}
	if tmp == n {
		return true
	}
	return false

}
