package LeetCode020

func isValid(s string) bool {
	//测试空字符串
	if s == "" {
		return true
	}

	//定义一个栈
	stack := make([]int32, len(s))
	stackLength := 0
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			//左括号,入栈
			stack[stackLength] = v
			stackLength++
		} else {
			//右括号,比较栈顶,匹配则移除,不匹配return false
			if stackLength == 0 {
				return false
			}
			if (v == ')' && stack[stackLength-1] == '(') || (v == ']' && stack[stackLength-1] == '[') || (v == '}' && stack[stackLength-1] == '{') {
				stackLength--
			} else {
				return false
			}
		}
	}
	return stackLength == 0 // 说明所有括号都出栈成功 匹配成功
}
