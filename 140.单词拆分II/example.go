package LeetCode140

import "strings"

// 这道题是上一题的升级版，既然需要返回所有结果，那么回溯基本上跑不了了。不过这道题不能强行回溯，会超时的，需要用到上一题的动态规划先判断字符串能不能被拆分，如果可以再进行回溯

func wordBreak(s string, wordDict []string) []string {
	wMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wMap[v] = true
	}

	//先通过DP判断字符串是否能被拆分
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	re := []string{}
	if !dp[len(s)] {
		return re
	}
	// 回溯走起
	var DFS = func(string, []string) {}
	DFS = func(ns string, r []string) {
		if len(ns) == 0 {
			re = append(re, strings.Join(r, " "))
			return
		}
		for i := 1; i <= len(ns); i++ {
			if wMap[ns[:i]] {
				DFS(ns[i:], append(r, ns[:i]))
			}
		}
	}
	DFS(s, []string{})
	return re
}
