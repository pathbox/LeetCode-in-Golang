package LeetCode1011

func shipWithinDays(weights []int, D int) int {
	// left 载重可能的最⼩值数组中的最大值 right 载重可能的最大值是数组元素之和
	left, right, mid := 0, 0, 0
	for _, v := range weights {
		right += v
		if v > left {
			left = v
		}
	}
	for left < right {
		mid = (left + right) / 2
		// 左边界搜索
		if canFinish(weights, mid, D) { // 满足条件收缩右边界
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 指定运载能力cap能否在D天内运完所有货物
func canFinish(weights []int, cap, D int) bool {
	daySum, day := 0, 1
	for i := 0; i < len(weights); {
		if daySum+weights[i] <= cap { // 能在容量内完成 day保持
			daySum += weights[i]
			i++
		} else {
			daySum, day = 0, day+1 // 不能在容量内完成 day需要多加1天
		}
	}
	return day <= D
}
