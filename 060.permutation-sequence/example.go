package main

import "fmt"

func main() {
	r := getPermutation(4, 17)
	fmt.Println(r)
}

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	// 存放结果字符,最后只输出一种排列结果，所以只要一个数组即可，不需要map或二维数组
	res := make([]byte, n)

	rec := make([]byte, n)
	for i := 0; i < n; i++ {
		rec[i] = byte(i) + '1'
	}

	// 由于排列的序号是从 1 开始的。
	// k 需要减去 1 ，好从 0 开始表示
	k--

	base := 1
	for i := 2; i < n; i++ {
		base *= i
	}
	for i := 0; i < n-1; i++ {
		idx := k / base
		fmt.Println("idx:", idx)
		res[i] = rec[idx]
		// 从 rec 中去除已经使用的数 rec[idx]
		rec = append(rec[:idx], rec[idx+1:]...)
		k %= base           //注意思考这里为何要取余，如果对这24个数以6个一组来分，那么k=16这个位置就是在第三组（k/6 = 2）中的第五个（k%6 = 4）
		base /= (n - i - 1) // (n-1)!
	}
	// 不要忘记最后一个数
	res[n-1] = rec[0]

	return string(res)
}
