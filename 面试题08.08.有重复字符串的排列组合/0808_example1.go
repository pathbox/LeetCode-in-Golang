package LeetCode0808

var result []string

func permutation(S string) []string {
	result = make([]string, 0)
	check := make(map[string]struct{}) // 用map来去重复
	used := make([]int, len(S))
	backtrace("", used, S, check)
	for i, _ := range check {
		result = append(result, i)
	}
	return result
}

func backtrace(cur string, used []int, S string, check map[string]struct{}) {
	if len(cur) == len(S) { // 长度相等 一种组合找到，加入到check
		check[cur] = struct{}{}
	}
	for i := 0; i < len(S); i++ { // 遍历每个字符
		if used[i] == 0 { // 这个字符没使用过
			used[i] += 1
			tmp := cur + string(S[i]) // 进行拼接
			backtrace(tmp, used, S, check)
			used[i] -= 1 // 回溯
		}
	}
}
