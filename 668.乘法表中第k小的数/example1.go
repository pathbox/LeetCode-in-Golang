package LeetCode668

// https://leetcode-cn.com/problems/kth-smallest-number-in-multiplication-table/solution/er-fen-cha-zhao-by-jason-2-5/

func findKthNumber(m int, n int, k int) int {
	l, r := 0, m*n
	for l < r {
		mid := l + (r-l)>>1
		cnt := 0 // 统计乘法表中小于等于mid的值的总数
		for i := 1; i <= m; i++ {
			if mid/i > n {
				cnt += n
			} else {
				cnt += mid / i
			}
		}
		if cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
