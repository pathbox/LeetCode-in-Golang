package LeetCode050

func myPow(x float64, n int) float64 {
	if n == 0 { // 递归终止条件
		return 1
	}

	if n < 0 { // // n为负数，先用其绝对值计算，然后取其倒数
		return 1 / myPow(x, -n)
	}
	if n%2 != 0 { // n 为奇数
		// n-1为偶数，再递归一次
		return x * myPow(x, n-1)
	}
	// 执行真正的运算逻辑，x的平方，n/2
	return myPow(x*x, n/2)
}
