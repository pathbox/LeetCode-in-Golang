package LeetCode1233

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	// 排序的目的是让父节点尽量在前面
	sort.Sort(sort.StringSlice(folder))
	out := make([]string, 0)
	out = append(out, folder[0])
	t := folder[0]
	for i := 1; i < len(folder); i++ {
		if strings.HasPrefix(folder[i], t+"/") == false { // 一定需要+"/"
			out = append(out, folder[i])
			t = folder[i] // 从 i开始，继续进行匹配,以floder[0]作为第一个前缀与后续字符比较，如果不是子集则写入输出的列表中，同时证明以floder[0]为前缀的所有字符串都过滤完成，此时将更新判断的前缀
		}
	}
	return out
}

// 时间复杂度O(nlogn) 比O(n^2)好
