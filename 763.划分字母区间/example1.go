package LeetCode763

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return []int{}
	}

	sitMap := make(map[string]int)
	runes := []rune(S)
	for index, v := range runes { //遍历得到每个字母最后位置
		sitMap[string(v)] = index
	}

	var curMaxSite int
	record, maxSite := -1, 0 //record 上一次完成计数位置，这里设为-1
	result := make([]int, 0)
	maxSite = sitMap[string(runes[0])]
	for i := 0; i < len(runes); {

		vString := string(runes[i])
		curMaxSite = sitMap[vString] // 当前字母最后面的位置

		if i == maxSite || i == len(runes)-1 { //当前位置等于记录最大的位置；当前位置为最后一个；
			result = append(result, i-record)
			record = i
			maxSite = 0

		} else if i == curMaxSite && curMaxSite > maxSite { //当前位置是该字母最后一次出现&&当前的位置大于记录最大位置eg:bbaasfc字符串当前位置i=3的情况
			result = append(result, i-record)
			record = i
			maxSite = 0
		} else if curMaxSite > maxSite && i != curMaxSite { //当前字母最后一次出现大于之前的最大记录&&该字母在后面还会出现
			maxSite = curMaxSite //变更最大位置记录
		}

		i++
	}
	return result
}
