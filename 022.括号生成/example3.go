package LeetCode022

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	helper(n, n, []byte{})
	return res
}

func helper(left, right int, tmp []byte) {
	if left == 0 && right == 0 {
		res = append(res, string(tmp))
		return
	}

	if left != 0 {
		tmp = append(tmp, '(')
		helper(left-1, right, tmp)
		tmp = tmp[:len(tmp)-1]
	}

	// 右括号数量不能大于左括号的，要不然合不回来，就不是合法数据了
	if right > left {
		tmp = append(tmp, ')')
		helper(left, right-1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}
