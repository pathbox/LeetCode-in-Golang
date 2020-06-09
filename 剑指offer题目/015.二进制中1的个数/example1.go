package Offer015

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}

	var res uint32
	for num > 0 {
		res += num & 1
		num >>= 1 // 左移一位
	}
	return int(res)
}

func hammingWeight2(num uint32) int {
	count := 0
	for 0 < num {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	return count
}
