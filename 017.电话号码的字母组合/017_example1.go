package LeetCode017

var num2String map[byte][]string = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if digits == "" {
		return res
	}

	for i := 0; i < len(digits); i++ {
		nums := num2String[digits[i]] // 得到每个号码对应的字符集合
		pre := res
		for _, num := range nums { // 取出每个字符
			for _, v := range pre { // 取出已经构造的每个元素，每个元素和现在遍历到的每个字符集合进行拼接
				res = append(res, v+num)
			}
		}
	}
	return res
}
