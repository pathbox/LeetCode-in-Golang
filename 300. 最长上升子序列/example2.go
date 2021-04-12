package LeetCode300

func lengthOfLIS(nums []int) int {
	d := []int{}
	for _, num := range nums {
		if len(d) == 0 || d[len(d)-1] < num {
			d = append(d, num)
		} else {
			l, r := 0, len(d)-1
			pos := r
			for l <= r {
				mid := (l + r) >> 1
				if d[mid] >= num {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			d[pos] = num
		}
	}
	return len(d)
}
