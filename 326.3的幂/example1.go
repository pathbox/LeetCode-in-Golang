package LeetCode326

func isPowerOfThree(n int) bool {
	if n > 1 {
		for n%3 == 0 {
			n /= 3
		}
	}
	return n == 1
}
