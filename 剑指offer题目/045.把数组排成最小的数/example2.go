package Offer045

import (
	"bytes"
	"sort"
	"strconv"
)

// Time: O(n*log(n)), Space: O(n)
func minNumber(nums []int) string {
	// 处理边界情况，如果数组为空或长度为0
	if nums == nil || len(nums) == 0 {
		return "" // 返回空的字符串
	}
	// 否则定义一个长度为n的字符串数组
	strs := make([]string, len(nums))
	// 遍历整数数组,把整数转成字符串,并存储到字符串数组
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	// 接下来对字符串进行排序,注意这里要自定义对比规则
	sort.Slice(strs, func(i, j int) bool {
		// 先把两种字符串的拼接结果写出来
		s12 := strs[i] + strs[j]
		s21 := strs[j] + strs[i]
		// 如果字符串s12小于s21，那么我们希望s1排在s2前面。
		return s12 < s21
	})
	// 把数组中的字符串依次拼接起来
	var buffer bytes.Buffer
	for i := range strs {
		buffer.WriteString(strs[i])
	}
	return buffer.String()
}
