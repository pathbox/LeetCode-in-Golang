package Offer056

func singleNumber(nums []int) int {
	h := map[int]int{}
	for _, v := range nums {
		h[v]++
	}
	for i, v := range h {
		if v == 1 {
			return i
		}
	}
	return 0
}
