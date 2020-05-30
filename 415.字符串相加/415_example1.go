package LeetCode415

import "fmt"

// 两个大数相加
func addStrings(num1 string, num2 string) string {
	i, j, carry, temp := len(num1)-1, len(num2)-1, 0, 0
	res := ""
	for i >= 0 || j >= 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}

		temp = n1 + n2 + carry // carry 是进位
		r := temp % 10         // 每次得到背刺和的个位上的数
		res = fmt.Sprintf("%d%s", r, res)
		carry = temp / 10 // 处理进位
		i--
		j--
	}
	if carry == 1 { // 说明最大位(最后一次相加)还有进位
		res = "1" + res
	}
	return res
}
