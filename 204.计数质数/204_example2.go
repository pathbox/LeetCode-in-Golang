package LeetCode204

import "math"

func countPrimes(n int) int {
	m := make([]int, n+1)
	count := 0

	for i := 0; i <= n; i++ {
		m[i] = 1
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrime(i) {
			for j := i * 2; j < n; j += i {
				m[j] = 0
			}
		}
	}
	for i := 2; i < n; i++ {
		if m[i] == 1 {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
