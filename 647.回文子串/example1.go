package LeetCode647

// https://leetcode-cn.com/problems/palindromic-substrings/solution/custerxue-xi-bi-ji-dong-tai-gui-hua-by-custer-go-4/

// 辅助函数 定义扩展函数，
// 输入是一个字符串，和两个向两边扩展的游标
// 输出是扩展过程中，回文子串的个数
func expand(s string, left, right int) int {
	count := 0 // 定义回文子串的个数为0
	// 当左右游标不超过左右边界并且它们指向的字符相等时
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] { // 左右游标限制条件
		count++
		left, right = left-1, right+1 // 两个游标向两边移动
	}
	return count // 循环结束后返回count即可
}

// 从中心向两边扩展回文子串的方法 Time: O(n^2), Space: O(1)
func countSubstrings(s string) int {
	// 边界情况,当字符串为空或长度为0时
	if s == "" || len(s) == 0 {
		return 0
	}
	// 初始回文子串的个数0
	count := 0
	for i := 0; i < len(s); i++ { // 遍历字符串
		// 以第i个字符为中心扩展子串
		// 得到的回文子串数量加到count上
		count += expand(s, i, i)
		count += expand(s, i, i+1)
	}
	return count
}
