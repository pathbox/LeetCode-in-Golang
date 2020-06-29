package LeetCode011

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	o := 0
	i, j := 0, len(height)-1
	for i != j {
		hi, hj := height[i], height[j]
		s := (j - i) * min(height[i], height[j])
		if s > o {
			o = s
		}

		if hi > hj { // 看哪边小，哪边移动，这样能保证更大的h一直有一边是
			j--
		} else {
			i++
		}
	}
	return o
}
