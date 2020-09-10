package LeetCode003

/*
核心：只增大不减小的滑动窗口
定义两个游标start和end，遍历字符串，end游标随着遍历无重复增大
若出现重复字符，则两个游标都增大index+1位（窗口大小不变，start游标滑动到重复位置后一位）
这时候由于窗口大小不变，窗口内可能存在重复，恰好遍历从start游标开始，如果没有重复，需要判断i+1与end的关系，超过的话，即i遍历到窗口之外，end再增大
遍历结束，end-start即为结果
*/

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		// 新的一个外循环表示左指针移动
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		// 内循环表示右指针不断移动
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
