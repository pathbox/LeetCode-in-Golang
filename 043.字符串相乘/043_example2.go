package LeetCode043

// 大数的相乘
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)
	res := make([]int, l1+l2)
	// 从两个数的最小位数开始比较
	for i := l1 - 1; i >= 0; i-- { // num1的每一位和num2的每一位进行相乘计算，得到结果存到res数组中.
		n1 := int(num1[i] - '0')
		for j := l2 - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			tmp := n1 * n2
			l, r := i+j, i+j+1  // 两个小于10的整数相乘，最大是81 占两位，l 表示整数位，r表示个数位
			sum := tmp + res[r] // 得到的值会先加个数位的，再和整数位的值相加
			res[r] = sum % 10   // 每次个数位都重新赋值
			res[l] += sum / 10  // 每次整数位进行累加,这次的整数位会变成下次的个数位累加
		}
	}

	for i, v := range res { // 会从高位开始遍历，找到不为0的值，说明当前i值就是最高位，得到res[i:]就得到结果
		if v != 0 {
			res = res[i:]
			break
		}
	}

	ans := ""
	for _, v := range res { // 把每个元素进行字符串的拼接
		ans += string(v + '0')
	}
	return ans

}
