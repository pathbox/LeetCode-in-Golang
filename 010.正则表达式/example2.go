package LeetCode010

// https://leetcode-cn.com/problems/regular-expression-matching/solution/golangdong-tai-gui-hua-by-bloodborne-5/

func isMatch(s string, p string) bool {
	lenStr, lenPattern := len(s), len(p)
	// 字符串s和正则p当前元素是否匹配
	matches := func(sIndex, pIndex int) bool {
		if sIndex < 1 {
			return false
		}
		if p[pIndex-1] == '.' || s[sIndex-1] == p[pIndex-1] {
			return true
		}
		return false
	}

	memo := make([][]bool, lenStr+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]bool, lenPattern+1)
	}
	memo[0][0] = true // 都是空字符串，当然可以
	for indexStr := 0; indexStr < lenStr; indexStr++ {
		for indexPattern := 1; indexPattern <= lenPattern; indexPattern++ {
			// 如果当前字符直接匹配，继承前面的结果
			if matches(indexStr, indexPattern) {
				memo[indexStr][indexPattern] = memo[indexStr-1][indexPattern-1]
				continue
			}
			if p[indexPattern-1] == '*' {
				// 要么是 ab / abc*，后面c*都是空字符这种匹配
				memo[indexStr][indexPattern] = memo[indexStr][indexPattern-2]
				if matches(indexStr, indexPattern-1) {
					// 要么是abbb / ab.* 或者abbb / abb* 这种匹配
					memo[indexStr][indexPattern] = memo[indexStr][indexPattern] || memo[indexStr-1][indexPattern]
				}
			}
		}
	}
	return memo[lenStr][lenPattern]
}
