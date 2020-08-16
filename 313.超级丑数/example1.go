package LeetCode313

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	uglyNums := make([]int, n)
	uglyNums[0] = 1

	indexes := make([]int, len(primes))
	i := 1
	for i < n {
		next := math.MaxInt32
		for j, index := range indexes {
			next = min(next, primes[j]*uglyNums[index])
		}
		for j, index := range indexes {
			if primes[j]*uglyNums[index] == next {
				indexes[j]++
			}
		}
		uglyNums[i] = next
		i++
	}
	return uglyNums[n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
