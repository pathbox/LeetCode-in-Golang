package LeetCode205

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map1 := make(map[string]string)
	map2 := make(map[string]string)
	for i := 0; i < len(s); i++ {
		if val, ok := map1[s[i:i+1]]; ok {
			if val != t[i:i+1] {
				return false
			}
		} else if val, ok := map2[t[i:i+1]]; ok {
			if val != s[i:i+1] {
				return false
			}
		} else {
			map1[s[i:i+1]] = t[i : i+1] // s->t 方向对应
			map2[t[i:i+1]] = s[i : i+1] // t->s 方向对应
		}
	}
	return true
}
