package LeetCode022

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(lRemain int, rRemain int, path string)

	dfs = func(lRemain int, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
			return
		}

		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		if lRemain < rRemain {
			dfs(lRemain, rRemain-1, path+")")
		}
	}

	dfs(n, n, "")
	return res
}
