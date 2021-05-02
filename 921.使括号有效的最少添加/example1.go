package LeetCode921

func minAddToMakeValid(S string) int {
	stL := make([]int, 0)
	stR := make([]int, 0)

	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stL = append(stL, i)
		} else if S[i] == ')' {
			if len(stL) > 0 {
				stL = stL[:len(stL)-1]
			} else {
				stR = append(stR, i)
			}
		}
	}
	return len(stL) + len(stR)
}

/*
使用两个栈存左括号索引和右括号索引，匹配到左括号入左括号栈，匹配到右括号，有特别操作，需要到左括号中查看是否已经有左括号可以与其构成合法括号组合，有的话，则进行“抵消操作”，将左括号出栈，而本身也不会入右括号栈。如果没有合适的左括号匹配，则入右括号栈
这样，最终在左括号栈和右括号栈就是没能合适组合成括号对的数量
*/
