package Offer014

func cuttingrope(n int) int {
	if n <= 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	result := 1
	// 以3为单位进行划分
	part := n / 3
	// 求解最后一段
	last := n % 3
	// 三种情况，清晰
	switch last {
	case 2:
		for i := 0; i < part; i++ {
			result *= 3
			result = result % 1000000007
		}
		result *= 2
		result = result % 1000000007
	case 1:
		for i := 0; i < part-1; i++ {
			result *= 3
			result = result % 1000000007
		}
		result *= 4 // 2*2
		result = result % 1000000007
	default:
		for i := 0; i < part; i++ {
			result *= 3
			result = result % 1000000007
		}
	}
	return result
}
