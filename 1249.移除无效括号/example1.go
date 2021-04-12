package LeetCode1249

func minRemoveToMakeValid(s string) string {
	//创建两个括号栈，分别存储左右括号的下标
	LStack := make([]int, 0)
	RStack := make([]int, 0)
	//括号删除栈，存储需要处理的括号的下标
	DelStack := make([]int, 0)

	for index, v := range s {
		if v == '(' {
			LStack = append(LStack, index)
		} else if v == ')' {
			if len(LStack) == 0 || len(LStack) <= len(RStack) {
				DelStack = append(DelStack, index)
			} else {
				RStack = append(RStack, index)
			}
		}
	}

	if len(LStack) > len(RStack) {
		DelStack = append(DelStack, LStack[len(RStack):]...)
	} else {
		DelStack = append(DelStack, RStack[len(LStack):]...)
	}
	offset := 0
	for _, value := range DelStack {
		s = s[:value-offset] + s[value+1-offset:]
		offset++
	}

	return s
}
