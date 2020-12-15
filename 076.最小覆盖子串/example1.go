package LeetCode076

import "math"

/*
算法：利用滑动窗口的思想。
https://leetcode-cn.com/problems/minimum-window-substring/solution/tong-su-qie-xiang-xi-de-miao-shu-hua-dong-chuang-k/
https://leetcode-cn.com/problems/minimum-window-substring/solution/si-lu-qing-xi-hao-li-jie-shuang-zhi-zhen-hua-dong-/

need保存t中出现的字符和次数，window保存窗口中满足t的字符和数量
初始化left，right指针
valid保存验证通过的字符数量
index，length保存字串的起始位置和长度
右指针右移更新window和valid
当valid == len（need）说明当前串儿满足要求，开始移动左指针，左指针的移动方式和逻辑与右指针相反
*/
func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	index, length := 0, math.MaxInt32

	// 分成两步维护window，满足need条件，往window中add，满足valid == len(need)了，left右移，寻找合适的index和length，当tempDel从window中减掉一次，破坏了 valid == len(need)时候，移动结束，继续向右
	for right < len(s) {
		// 1. 寻找满足包含t的滑动窗口
		tempAdd := s[right]
		right++
		if _, ok := need[tempAdd]; ok { // 在need中
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				valid++
			}
		}

		// 2. 滑动窗口操作
		for valid == len(need) { // 窗口中，包含了所有need
			if right-left < length {
				index = left // 重置起始位置
				length = right - left
			}
			tempDel := s[left] // 要被缩小的left
			left++             // 左指针移动

			if _, ok := need[tempDel]; ok {
				window[tempDel]--
				if window[tempDel] < need[tempDel] {
					valid--
				}
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[index : index+length]
}
