package LeetCode786

func kthSmallestPrimeFraction(arr []int, k int) []int {
	var low float64 = 0
	var high float64 = 1
	ans := []int{0, 1}
	for high-low > 1e-9 {
		mid := low + (high-low)/2.0
		res := under(mid, arr)
		if res[0] < k { // 说明数量太少
			low = mid
		} else {
			ans[0] = res[1]
			ans[1] = res[2]
			high = mid
		}
	}
	return ans
}

func under(x float64, arr []int) []int {
	numer, denom, count, i := 0, 1, 0, -1
	for j := 1; j < len(arr); j++ {
		for float64(arr[i+1]) < float64(arr[j])*x { // primes[i] / primes[j] < x
			i++
		}
		count += i + 1
		if i >= 0 && numer*arr[j] < denom*arr[i] {
			numer = arr[i]
			denom = arr[j]
		}
	}
	return []int{count, numer, denom}
}

/*
时间复杂度：O(N \log W)O(NlogW)，其中 NN 是 primes 的长度，W 是二分查找的区间宽度，(hi - lo) / 1e-9 等于 10^9

空间复杂度：O(1)
*/
