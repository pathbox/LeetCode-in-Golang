package LeetCode091

// 和爬楼梯思路类似，但是有很多条件限制。每次可以取一个字符，或取两个字符等等

// 用变量cur表示当前字符串长度的最多译码方法数量，由于每次遍历cur的取值之和前两次有关，所以每次遍历开始的时候cur,past分别表示上一次和上上一次动态规划值。没有必要申请切片或者数组，这样空间复杂度就位O(1).
// 注意特殊情况就是当字符串第一个字母为0时，直接返回0，因为无法解码这种情况

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur, tmp := 1, 1, 0
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2':
			return 0
		case s[i] == '0':
			cur = pre
		case (s[i] <= '6' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1'):
			tmp = cur
			cur = cur + pre
			pre = tmp
		default:
			pre = cur
		}
	}
	return cur
}
