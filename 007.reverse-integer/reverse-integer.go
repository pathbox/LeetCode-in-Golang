package Problem007

import "math"

func reverse(x int) int {
	sign := 1 // sign 的1 -1 表示这个整数的符号

	// 处理负数
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0    // 结果值
	for x > 0 { //不断的取余数,取整方法，得到x的末尾数字
		// 取出x的末尾
		temp := x % 10
		// 放入 res 的开头
		res = res*10 + temp
		// x 去除末尾
		x = x / 10
	}

	res = sign * res // 加上符号

	// 处理 res 的溢出问题
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res

}
