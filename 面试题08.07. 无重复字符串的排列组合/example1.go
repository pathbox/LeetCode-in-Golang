package LeetCode0807

var result []string

func permutation(S string) []string {
	result = make([]string, 0)
	used := make([]int, len(S)) // 使用used来记录该位置是否选过，来过滤可选列表
	backTrace("", S, used)
	return result
}

func backTrace(cur, S string, used []int) {
	if len(cur) == len(S) {
		result = append(result, cur)
		return
	}

	for i := 0; i < len(S); i++ {
		if used[i] == 0 {
			used[i]++
			tmp := cur + string(S[i])
			backTrace(tmp, S, used)
			used[i]-- // 可选列表的回溯
		}
	}
}
