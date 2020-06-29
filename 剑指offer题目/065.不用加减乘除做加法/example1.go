package Offer065

// https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/solution/mei-ri-suan-fa-day-66-jing-dian-mian-shi-ti-bu-yon/

// 异或做加法,与后左移一位做进位
func add(a, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1 // 异或操作, // 与操作后左移一位
	}
	return a
}
