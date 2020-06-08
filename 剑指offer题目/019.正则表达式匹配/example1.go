package Offer019

// 递归
// 终止条件：p为空时，s为空返回true，s不空返回false
// first标志首位是否相等，p[0]是.也表示相等
// (1)如果p的第2位为*
//    ①首位相同 则等同于判断s[1:]和p （忽略s[0]）
//    ②首位不同 则等同于判断s和p[2:] （将p前两位x*和空值配对）
// (2)否则不要期待奇迹出现了，
//    首位如果相等就各向后移一位，判断s[1:]和p[1:]
//    如果不等就是false
// 这里利用了go字符串的截断符号

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	var first bool                                   // 可以理解为首位，也可以理解为当前开始的匹配位
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') { // 递归的一种条件情况
		first = true
	}
	if len(p) > 1 && p[1] == '*' { // 第二种情况
		return (first && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}
	return first && isMatch(s[1:], p[1:]) // 一般情况，继续往下递归
}
