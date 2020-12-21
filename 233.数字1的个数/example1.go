package LeetCode233

// https://leetcode-cn.com/problems/number-of-digit-one/solution/custerxue-xi-bi-ji-shu-xue-by-custergo/

// 暴力解法 Time: O(n*log10(n)), Space: O(1)
func countDigitOne(n int) int {
	// 边界情况
	if n < 1 {
		return 0
	}
	count := 0                // 初始化计数器为0
	for i := 1; i <= n; i++ { // 接着i从整数1遍历到整数n
		num := i       // 将整数i赋值为num
		for num != 0 { // 当num不等于0时不断执行以下操作
			if num%10 == 1 { // 先对10取模看个数上的数字是否等于1
				count++ // 如果是计数器+1
			}
			num = num / 10 // 然后除以10来更新num
		}
	}
	return count // 循环结束后返回count即可
}

// 使用数学方法 Time: O(log10(n)), Space: O(1)
func countDigitOne(n int) int {
	// 边界情况
	if n < 1 {
		return 0
	}
	count, factor := 0, 1 // 初始化计数器为0,和位数初始化为1
	for n/factor != 0 {   // 当整数n除以factor不等于0时
		// 不断执行之下操作
		// 先求出当前位上的数字digit
		digit := (n / factor) % 10 // n除以factor再对10取模
		// 然后计算更高位的数字high
		high := n / (10 * factor) // n除以10倍的factor
		if digit == 0 {           // 如果当前位数的数字等于0
			count += high * factor // 计数器则加上high乘以factor
		} else if digit == 1 { // 如果当前位数字等于1
			count += high * factor    // 计数器不仅要加上igh乘以factor
			count += (n % factor) + 1 // 还要加上低位数字(n对factor取模即可求出)再加1
		} else { // 如果是其他情况
			count += (high + 1) * factor // 计数器则加上(high+1)乘以factor
		}
		// 计算完当前factor位上1出现的次数
		factor = factor * 10 // factor乘以10进行更新,准备计算下一个十进制位
	}
	return count // 循环结束后返回count即可。
}
