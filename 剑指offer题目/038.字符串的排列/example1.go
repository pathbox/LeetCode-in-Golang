package Offer038

func permutation(s string) []string {
	res := make(map[string]bool)
	c := []byte(s)
	dfs(res, c, 0)
	RES := []string{}
	for key := range res {
		RES = append(RES, key)
	}
	return RES
}

func dfs(res map[string]bool, c []byte, i int) {
	if i == len(c)-1 { // 满足一种条件了
		res[string(c)] = true
		return
	} else {
		for j := i; j < len(c); j++ { // j := i
			c[i], c[j] = c[j], c[i] // 交换
			dfs(res, c, i+1)
			c[i], c[j] = c[j], c[i] // 回溯 c恢复为当次传入时候的值
		}
	}
}
