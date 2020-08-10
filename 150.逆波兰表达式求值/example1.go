package LeetCode150

import "strconv"

func evalRPN(tokens []string) int {
	tlen := len(tokens)
	stack := []int{}
	for i := 0; i < tlen; i++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			if s, err := strconv.Atoi(tokens[i]); err == nil {
				stack = append(stack, s)
			}
			continue
		}

		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		switch tokens[i] {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, a-b)
		case "*":
			stack = append(stack, a*b)
		case "/":
			stack = append(stack, a/b)
		}
		//fmt.Println(stack)
	}
	return stack[len(stack)-1]
}
