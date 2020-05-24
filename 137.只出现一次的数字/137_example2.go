package LeetCode137

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		a = (a ^ num) & (^b) // num先对b按位清除；再与a异或
		b = (b ^ num) & (^a) // num先对b按位清除；再与a异或
	}
	return a
}
