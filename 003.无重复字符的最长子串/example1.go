package LeetCode003

/*
核心：只增大不减小的滑动窗口
定义两个游标start和end，遍历字符串，end游标随着遍历无重复增大
若出现重复字符，则两个游标都增大index+1位（窗口大小不变，start游标滑动到重复位置后一位）
这时候由于窗口大小不变，窗口内可能存在重复，恰好遍历从start游标开始，如果没有重复，需要判断i+1与end的关系，超过的话，即i遍历到窗口之外，end再增大
遍历结束，end-start即为结果
*/

func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先设置所有的字符都没有见过
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 在s[left:i+1]中去除当前重复的s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
	}
	return maxLen
}
