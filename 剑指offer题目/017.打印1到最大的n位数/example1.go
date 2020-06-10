package Offer017

func printNumbers(n int) []int {
	max := getMax(n)
	res := make([]int, 0)
	for i := 1; i <= max; i++ {
		res = append(res, i)
	}
	return res
}

func getMax(n int) int {
	max := 1
	for i := 0; i < n; i++ {
		max *= 10
	}
	return max - 1
}
