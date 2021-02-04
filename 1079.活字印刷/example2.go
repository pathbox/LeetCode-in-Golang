package LeetCode1079

func numTilePossibilities(tiles string) int {
	count := [26]int{}
	for i := 0; i < len(tiles); i++ {
		count[tiles[i]-'A']++
	}
	return dfs(count)
}

func dfs(count [26]int) int {
	res := 0
	for i := 0; i < 26; i++ {
		if count[i] == 0 {
			continue
		}
		res++
		count[i]--
		res += dfs(count)
		count[i]++ // 只需要重置字符频数数组
	}
	return res
}
