package Offer056

func singleNumber(nums []int) int {
	a := 0
	b := 0
	for _, c := range nums {
		a, b = a&^c|b&c, b&^c|^a&^b&c
	}
	return ^a & b
}
