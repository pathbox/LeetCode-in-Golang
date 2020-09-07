package Offer055

func singleNumbers(nums []int) []int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	lb := res & (-res)
	val := make([]int, 2)
	for _, n := range nums {
		if lb&n == 0 {
			val[0] ^= n
		} else {
			val[1] ^= n
		}
	}
	return val
}
