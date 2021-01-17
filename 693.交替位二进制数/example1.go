package LeetCode693

func hasAlternatingBits(n int) bool {
	// 定义：n 的最后一位是0，则flag为true，否则为false
	var flag bool
	if (n & 1) == 0 {
		flag = true // 表示最后一位是0
	}

	n = n >> 1

	// 从倒数第二位开始
	for n != 0 {
		// 如果此时n的最后一位是0，并且上一次也是0的话，那么就直接返回false
		if (n&1) == 0 && flag {
			return false
		}

		// 如果此时n的最后一位是1，并且上一次也是1的话，那么就直接返回false
		if (n&1) == 1 && !flag {
			return false
		}

		// 如果没有命中前两种情况，就说明和上一次的最后一位的不同
		n = n >> 1
		flag = !flag
	}
	return true

}
