package LeetCode093

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := new([]string)
	dfs(s, ans, []string{})
	return *ans
}

func dfs(s string, ans *[]string, old []string) {
	if len(old) == 3 {
		if isIp(s) {
			old = append(old, s)
			*ans = append(*ans, strings.Join(old, "."))
		}
		return
	}
	for i := 1; i <= 3 && i < len(s); i++ {
		if isIp(s[:i]) {
			dfs(s[i:], ans, append(append([]string{}, old...), s[:i]))
		}
	}
	return
}

func isIp(s string) bool {

	if len(s) == 0 {
		return false
	}

	if len(s) > 3 {
		return false
	}

	if len(s) > 1 && s[0] == '0' {
		return false
	}

	v, _ := strconv.Atoi(s)
	return v <= 255
}
