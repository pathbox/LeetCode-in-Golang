package LeetCode077

func combine(n int, k int) [][]int {
	if k > n {
		return nil
	}
	var res [][]int
	var arr []int
	if n == 0 {
		return [][]int{{}}
	}
	getCombine(1, k, &arr, &res, n)
	return res
}

func getCombine(n int, k int, arr *[]int, res *[][]int, max int) {
	if len(*arr) == k {
		temp := make([]int, len(*arr))
		copy(temp, *arr)
		*res = append(*res, temp)
		return
	}

	if k-len(*arr) > max-n+1 { // 如果还需要取的数量大于 还能够取的数量，则没有必要再进行了，因为剩下的数量达不到取的数量了
		return
	}

	// getCombine(n+1, k, arr, res, max)
	*arr = append(*arr, n)
	getCombine(n+1, k, arr, res, max)
	res2 := *arr
	res2 = res2[:len(res2)-1]
	*arr = res2
}

/*
二：如果始终达不到k呢？比如当n为8，k为5的时候，我们发现，如果DFS到数组[5,6]了，那么就算把剩下7、8全部填进去，也不能满足数组长度等于5的条件。那么这里剪枝的规律就是：k-len(*arr)>max-n+1，max就是8，n表示当前填的是第几位

*/
