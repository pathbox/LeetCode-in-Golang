package LeetCode1518

func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for {
		a := numBottles / numExchange
		b := numBottles % numExchange
		numBottles = a + b
		if a > 0 {
			res += a
		}
		if a == 0 {
			return res
		}

	}
	return res
}
