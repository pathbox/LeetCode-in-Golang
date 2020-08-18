package LeetCode454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	count := make(map[int]int, 0)
	for _, v1 := range A {
		for _, v2 := range B {
			count[v1+v2]++
		}
	}
	result := 0
	for _, v1 := range C {
		for _, v2 := range D {
			result += count[-(v1 + v2)]
		}
	}
	return result
}
