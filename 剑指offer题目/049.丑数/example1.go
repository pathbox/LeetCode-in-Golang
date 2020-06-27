package Offer049

// https://leetcode-cn.com/problems/chou-shu-lcof/solution/chou-shu-ii-qing-xi-de-tui-dao-si-lu-by-mrsate/
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	result := []int{1}
	res := 0
	i, j, k := 0, 0, 0 // 三个指针
	// 本质是三个数组的排序合并
	for n > 1 {
		x := result[i] * 2 // 三个数值
		y := result[j] * 3
		z := result[k] * 5

		res = min(min(x, y), z) // 得到当前最小的值
		// 如果最小值是该数值，该数值的索引走下一个，要不然索引保持不变等待下一轮
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
	return result[len(result)-1] // 返回最后一个数
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
