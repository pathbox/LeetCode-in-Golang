package LeetCode204

func countPrimes(n int) int {
	a := make([]bool, n)
	cnt := 0
	for i := 2; i < n; i++ {
		if a[i] {
			continue
		}
		for j := i + 1; j < n; j += i { // 所有 i 的倍数都置 true，因为 i 是他们的因子，肯定不是质数了,// 在n范围内，以小数开始，得到他们的倍数值，这些值都不是质数
			a[j] = true
		}
		cnt++
	}
	return cnt
}
