package LeetCode227

import "strconv"

/*
1、先给第一个数字加一个默认符号+，变成+1-12+3。
2、把一个运算符和数字组合成一对儿，也就是三对儿+1，-12，+3，把它们转化成数字，然后放到一个栈中。
3、将栈中所有的数字求和，就是原算式的结果
*/
func calculate(s string) int {
	numstr := "0"
	stack := []int{0}
	op := []byte{'+'}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numstr += string(s[i])
		}
		if i == len(s)-1 || '+' == s[i] || '-' == s[i] || '*' == s[i] || '/' == s[i] {
			num, err := strconv.Atoi(numstr)
			if err != nil {
				break
			}
			// 乘除法的优先级体现在可以和devidor计算结合，+ - 只能把devidor再放回栈中，然后再把自己放到栈中
			// 乘除法优先于加减法体现在，乘除法可以和栈顶的数结合，而加减法只能把自己放入栈
			devidor := stack[len(stack)-1] // 栈最顶层树
			stack = stack[:len(stack)-1]
			if '*' == op[len(op)-1] {
				stack = append(stack, devidor*num)
			} else if '/' == op[len(op)-1] {
				stack = append(stack, devidor/num)
			} else if '+' == op[len(op)-1] {
				stack = append(stack, devidor)
				stack = append(stack, num)
			} else if '-' == op[len(op)-1] {
				stack = append(stack, devidor)
				stack = append(stack, -num)
			}
			op = op[:len(op)-1]
			op = append(op, s[i])
			numstr = ""
		}
	}

	res := 0
	for _, v := range stack { // 最后遍历stack 把所有元素累加
		res += v
	}
	return res
}
