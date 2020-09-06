package Offer038

func permutation(s string) []string {
	var ret []string
	var m = make(map[string]struct{})
	bfs(s, ``, m)

	for key := range m {
		ret = append(ret, key)
	}

	return ret
}

func bfs(remained string, now string, m map[string]struct{}) {
	if len(remained) == 0 {
		m[now] = struct{}{}
		return
	}
	for i := 0; i < len(remained); i++ {
		s := remained[0:i] + remained[i+1:]
		bfs(s, now+string(remained[i]), m)
	}
}
