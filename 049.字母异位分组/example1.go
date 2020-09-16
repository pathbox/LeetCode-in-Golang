package LeetCode049

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	group := make(map[string]int)
	result := make([][]string, 0)
	index := 0
	for _, v := range strs {
		vArr := strings.Split(v, "")
		sort.Strings(vArr)
		toStr := strings.Join(vArr, "")
		if idx, ok := group[toStr]; !ok {
			group[toStr] = index
			result = append(result, []string{v})
			index++
		} else {
			result[idx] = append(result[idx], v)
		}
	}
	return result
}
