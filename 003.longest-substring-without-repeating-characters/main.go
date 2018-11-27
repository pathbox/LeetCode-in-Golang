package main

import "fmt"

func main() {
	s := "abcabcwwssssssssssssssssss"
	l := lengthOfLongestSubstringMap(s)
	fmt.Println(l)
}

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
			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}
	return maxLen
}

func lengthOfLongestSubstringMap(s string) int {
	length := 0
	max := 0
	start := 0

	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		fmt.Println(m)
		if index, ok := m[s[i]]; ok {
			if index < start {
				length++
			} else {
				start = index + 1
				length = i - index
			}

		} else {
			length++
		}

		m[s[i]] = i // 已经遍历过第i位的字符，存到m中

		if length > max {
			max = length
		}
	}

	return max
}
