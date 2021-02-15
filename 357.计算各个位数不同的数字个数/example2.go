package LeetCode357

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	used := make([]bool, 10)
	return dfs(n, 0, used)
}

func dfs(n, index int, used []bool) int {
	if index == n {
		return 0
	}
	result := 0
	for num := 0; num <= 9; num++ {
		// 剪枝：多位数时，第一个数字不能为0
		if n >= 2 && index == 1 && num == 0 {
			continue
		}
		if used[num] {
			continue
		}
		result += dfs(n, index+1, used) + 1
		used[num] = false
	}
	return result
}
