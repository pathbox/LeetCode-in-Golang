package main

import "fmt"

func main() {
	n := 10
	r := climbStairs(n)
	fmt.Println("Result:", r)
}

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return 1
	}

	r := climbStairs(n-1) + climbStairs(n-2)
	return r
}
