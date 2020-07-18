package LeetCode344

// 首尾交换即可

func reverseString(s []byte) {
	n := len(s)
	mid := n / 2
	for i := 0; i < mid; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}
