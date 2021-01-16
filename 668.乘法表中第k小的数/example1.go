package LeetCode668

// https://leetcode-cn.com/problems/kth-smallest-number-in-multiplication-table/solution/er-fen-cha-zhao-by-jason-2-5/
// https://leetcode-cn.com/problems/kth-smallest-number-in-multiplication-table/solution/can-kao-liao-ping-lun-qu-da-lao-de-fang-fa-zhong-x/
// 第k小的数不是索引位置在乘法表的第k位，乘法表中会有很多数出现在不同的位置
// 第i行是i的倍数，所以mid/i表示mid可以分成多少个i的倍数，n是表示这一行列的数量。如果mid/i > n表示mid可以分成多少个i的倍数大于n列数，说明这一列的所有数都小于mid/i。
// 换一种方式理解，抓住每一行i是i的倍数。mid/i表示mid可以分成多少个i的倍数，也就是mid在第i行的第几位，这个位数可能大于n也可能小于n
func findKthNumber(m int, n int, k int) int {
	l, r := 0, m*n
	for l < r
		mid := l + (r-l)>>1
		cnt := 0 // 统计乘法表中小于等于mid的值的总数
		for i := 1; i <= m; i++ {
			// mid/i 当前行mid值所在的位置
			if mid/i > n {
				cnt += n
			} else {
				cnt += mid / i
			}
		}
		if cnt >= k {
			r = mid
		} else {
			l = mid + 1 // 说明当前mid这个值在目标值的左边（比目标值小），所以可以缩小边界
		}
	}
	return l
}
