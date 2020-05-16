package LeetCode216

var result [][]int

func combinationSum3(k int, n int) [][]int {
	result = make([][]int, 0)
	cur := make([]int, 0)
	backtrace(1, cur, 0, k, n)
	return result
}

func backtrace(start int, cur []int, count, k, n int) {
	if count == n && len(cur) == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		result = append(result, tmp)
		return
	}
	for i := start; i <= 9; i++ {
		cur = append(cur, i)
		count += i
		backtrace(i+1, cur, count, k, n)
		count -= i
		cur = cur[:len(cur)-1]
	}
}
