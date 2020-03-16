package LeetCode007

func reverse(x int) int {
	y := 0
	for x != 0 { // 0 说明所有位的数反转过了
		y = y*10 + x%10                         // %10 取最后一个位的数
		if !(-(1<<31) <= y && y <= (1<<31)-1) { // y不在这个范围 返回0
			return 0
		}
		x /= 10 // 最后一位的数取过了，去除最后一位的数
	}
	return y
}
