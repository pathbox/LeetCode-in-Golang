package LeetCode1208

func equalSubstring(s string, t string, maxCost int) int {
	if len(s) == 0 {
		return 0
	}

	left, right := 0, 0
	sum, result := 0, 0
	for left <= right && right <= len(s)-1 {
		sum += abs(int(s[right]) - int(t[right]))
		// 如果当前花销大于预算，则left移动，并且去除sum中left的花销
		for sum > maxCost {
			sum -= abs(int(s[left]) - int(t[left]))
			left++
			if left > len(s)-1 {
				break
			}
		}

		result = max(result, right-left+1) // 每轮都会计算一次result  长度就是 right-left+1
		right++
	}

	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
