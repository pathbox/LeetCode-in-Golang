package LeetCode371

/*
a + b 的问题拆分为 (a 和 b 的无进位结果) + (a 和 b 的进位结果)
无进位加法使用异或运算计算得出
进位结果使用与运算和移位运算计算得出
循环此过程，直到进位为 0

https://leetcode-cn.com/problems/sum-of-two-integers/solution/0msfu-xian-ji-suan-ji-zui-ji-ben-de-jia-fa-cao-zuo/
*/
func getSum(a int, b int) int {
	and := (a & b) << 1 // 进位
	xor := a ^ b        // 低位
	for and != 0 {      // 进位为0则说明加法不再继续
		b = and
		a = xor
		and = (a & b) << 1
		xor = a ^ b
	}
	return xor
}
