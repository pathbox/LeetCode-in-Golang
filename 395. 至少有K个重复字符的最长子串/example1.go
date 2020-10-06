package LeetCode395

/*
1、找到所有总计数小于k的字母，这些字母铁定不能作为最终结果返回的
2、因此把他们作为分隔符split用，把字符串s分成多个短字符串进行求解
3、把求得的最大长度放在result里返回
4、在分解字符串的时候，如果子字符串长度小于等于result，那就可以无视了

*/
func longestSubstring(s string, k int) int {
	charNum := make(map[rune]int)
	for _, v := range s {
		charNum[v]++
	}
	var split []int
	// 找到所有不满足条件的字符作为分割符，进行分治
	for index, v := range s {
		if charNum[v] < k {
			split = append(split, index)
		}
	}
	if len(split) == 0 { // 没有不满足条件的
		return len(s)
	}

	split = append(split, len(s)) // 把最后的右边加进去
	result, left := 0, 0
	for _, v := range split {
		// 总长度不超过已知最长的就没必要比较了
		if v-left+1 <= result {
			continue
		}
		result = max(result, longestSubstring(s[left:v], k))
		left = v + 1
	}
	return result
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
