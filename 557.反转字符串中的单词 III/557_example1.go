package LeetCode557

func reverseWords(s string) string {
	bt := []byte(s)
	if len(bt) == 0 {
		return s
	}
	slow := 0
	for i := 0; i <= len(bt)-1; i++ {
		if bt[i] == ' ' {
			fast := i - 1
			for slow < fast {
				bt[slow], bt[fast] = bt[fast], bt[slow]
				fast--
				slow++
			}
			slow = i + 1
		}

		if i == len(bt)-1 {
			fast := i
			for slow < fast {
				bt[slow], bt[fast] = bt[fast], bt[slow]

				fast--
				slow++
			}
		}
	}
	return string(bt)
}
