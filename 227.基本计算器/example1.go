package LeetCode227

/*
1、先给第一个数字加一个默认符号+，变成+1-12+3。
2、把一个运算符和数字组合成一对儿，也就是三对儿+1，-12，+3，把它们转化成数字，然后放到一个栈中。
3、将栈中所有的数字求和，就是原算式的结果
*/
func calculate(s string) int {
	stack := []int{}
	preSign := '+'
	num := 0
	ans := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}

		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return ans
}
