package LeetCode093

import "strings"

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	if len(s) < 4 || len(s) > 12 {
		return res
	}
	dfs(make([]string, 0), s, &res)
	return res
}

func dfs(ss []string, s string, res *[]string) {
	lens := len(ss)
	if lens == 4 {
		if len(s) == 0 { // 表示整个字符串被dfs过
			*res = append(*res, ss[0]+"."+ss[1]+"."+ss[2]+"."+ss[3])
		}
		return
	}
	for i := 1; i <= 3; i++ { // 每个元素最大长度是3
		if len(s) < i {
			return
		}
		sss := s[0:i]
		strLen := len(sss)
		if strLen == 3 && strings.Compare(sss, "255") > 0 {
			return
		}
		if strLen > 1 && s[0] == '0' {
			return
		}
		ss = append(ss, sss)
		dfs(ss, s[i:], res)
		ss = ss[:len(ss)-1] // 回溯
	}
}
