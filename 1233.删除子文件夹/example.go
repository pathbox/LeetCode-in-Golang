package LeetCode1233

import (
	"sort"
	"strings"
)

/*
1.排序
2. 初始父目录
3. +"/" 与之后前缀匹配
4. 匹配失败，后一位是父目录，进入返回集合，并且后面的所有和当前也肯定会失败，以后一位为新的父目录与之后进行匹配
*/

func removeSubfolders(folder []string) []string {
	// 排序的目的是让父节点尽量在前面
	sort.Sort(sort.StringSlice(folder))
	out := make([]string, 0)
	out = append(out, folder[0])
	t := folder[0]
	for i := 1; i < len(folder); i++ {
		if strings.HasPrefix(folder[i], t+"/") == false { // 一定需要+"/" folder[i]不是子目录,加入返回集合
			out = append(out, folder[i]) // 匹配失败，说明 folder[i]是父目录
			t = folder[i]                // 从 i开始，继续进行匹配,以floder[0]作为第一个前缀与后续字符比较，如果不是子集则写入输出的列表中，同时证明以floder[0]为前缀的所有字符串都过滤完成，此时将更新判断的前缀
			// 为什么可以这样呢？因为前面的才会是父目录，如果前一位和后一位没有相同的前缀，说明后一位也是一个父目录，加入到返回集合后，再以后一位这个父目录，作为父目录的匹配标准，和后面的进行匹配，看后面的是否是这个的子目录
			// 所以不需要双重循环
		}
	}
	return out
}

// 时间复杂度O(nlogn) 比O(n^2)好
